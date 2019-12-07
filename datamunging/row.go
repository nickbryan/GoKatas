package datamunging

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

// Row represents a single row in the data.
type Row struct {
	Id   string
	A, B float64
}

// Spread calculates the difference between A and B to give the Spread.
func (r Row) Spread() float64 {
	return math.Abs(r.A - r.B)
}

// Rows represents a set of data.
type Rows []Row

// ReadFrom scans the given io.Reader and extracts Id, A and B into a Rows. Each Row represents a single
// parsed line in the data. The first two lines (header and separator) are skipped.
func (rs *Rows) ReadFrom(r io.Reader, colId, colA, colB string) error {
	rp := strings.NewReplacer("-", "", "*", "")

	ls := bufio.NewScanner(r)
	ls.Split(bufio.ScanLines)

	i, a, b, err := readHeaderLine(ls, colId, colA, colB)
	if err == io.EOF {
		return nil
	}

	if err != nil {
		return fmt.Errorf("unable to scan header line: %w", err)
	}

	var line, col int
	for ls.Scan() {
		var row Row

		cs := bufio.NewScanner(strings.NewReader(rp.Replace(ls.Text())))
		cs.Split(bufio.ScanWords)

		col = 0
		for cs.Scan() {
			switch col {
			case i:
				row.Id = cs.Text()
			case a:
				v, err := strconv.ParseFloat(cs.Text(), 64)
				if err != nil {
					return fmt.Errorf("unabled to parse %s as float: %w", colA, err)
				}

				row.A = v
			case b:
				v, err := strconv.ParseFloat(cs.Text(), 64)
				if err != nil {
					return fmt.Errorf("unabled to parse %s as float: %w", colB, err)
				}

				row.B = v
			}
			col++
		}

		if col != 0 {
			*rs = append(*rs, row)
		}
	}

	if err := ls.Err(); err != nil {
		return fmt.Errorf("unable to scan line %d: %w", line, err)
	}

	return nil
}

func readHeaderLine(s *bufio.Scanner, i, a, b string) (ii, ai, bi int, err error) {
	var scanned int
	ii, ai, bi = -1, -1, -1

	s.Scan()
	if s.Text() == "" {
		return ii, ai, bi, io.EOF
	}

	rs := bufio.NewScanner(strings.NewReader(s.Text()))
	rs.Split(bufio.ScanWords)

	for rs.Scan() {
		switch rs.Text() {
		case i:
			ii = scanned
		case a:
			ai = scanned
		case b:
			bi = scanned
		}
		scanned++
	}

	if err = rs.Err(); err != nil {
		err = fmt.Errorf("unable to scan header column %d: %w", scanned+1, err)
	}

	if ii == -1 || ai == -1 || bi == -1 {
		err = fmt.Errorf("unable to read header: %s = %d, %s = %d, %s = %d", i, ii, a, ai, b, bi)
	}

	return ii, ai, bi, err
}

// MinSpread returns the Row with the minimum Spread value in the set.
func (rs *Rows) MinSpread() Row {
	var msr Row

	for i, r := range *rs {
		if i == 0 || r.Spread() <= msr.Spread() {
			msr = r
		}
	}

	return msr
}
