package datamunging

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

// Row represents a single row in the data.
type Row struct {
	Id   string
	A, B float64
}

// Spread calculates the difference between A and B to give the Spread.
func (r *Row) Spread() float64 {
	return math.Abs(r.A - r.B)
}

// Rows represents a set of data.
type Rows []*Row

// ReadFrom scans the given io.Reader and extracts Id, A and B into a Rows. Each Row represents a single
// parsed line in the data. The first two lines (header and separator) are skipped.
func (rs *Rows) ReadFrom(r io.Reader) error {
	ls := bufio.NewScanner(r)
	ls.Split(bufio.ScanLines)

	var scanned int
	for ls.Scan() {
		var row Row

		// skip the first two lines of the file (header and blank line before data)
		if scanned++; scanned <= 2 {
			continue
		}

		if _, err := fmt.Fscan(strings.NewReader(ls.Text()), &row.Id, &row.A, &row.B); err != nil {
			return fmt.Errorf("unable to read values from line %d: %w", scanned, err)
		}

		*rs = append(*rs, &row)
	}

	if err := ls.Err(); err != nil {
		return fmt.Errorf("unable to scan line %d: %w", scanned, err)
	}

	return nil
}

// MinSpread returns the Row with the minimum Spread value in the set.
func (rs *Rows) MinSpread() *Row {
	var msr *Row

	for _, r := range *rs {
		if msr == nil || r.Spread() <= msr.Spread() {
			msr = r
		}
	}

	if msr == nil {
		return &Row{}
	}

	return msr
}
