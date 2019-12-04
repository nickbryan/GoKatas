package datamunging

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Row struct {
	Day      uint
	Min, Max float64
}

type Rows []Row

func Parse(r io.Reader) (Rows, error) {
	ls := bufio.NewScanner(r)
	ls.Split(bufio.ScanLines)

	var rows Rows
	var scanned int
	for ls.Scan() {
		// skip the first two lines of the file (header and blank line before data)
		if scanned++; scanned <= 2 {
			continue
		}

		row := Row{}
		if _, err := fmt.Fscan(strings.NewReader(ls.Text()), &row.Day, &row.Min, &row.Max); err != nil {
			return nil, fmt.Errorf("unable to read values from line %d: %w", scanned, err)
		}

		rows = append(rows, row)
	}

	if err := ls.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan line: %w", err)
	}

	return rows, nil
}
