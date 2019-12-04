package datamunging

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Parse(r io.Reader) (Rows, error) {
	ls := bufio.NewScanner(r)
	ls.Split(bufio.ScanLines)

	var rows Rows
	var scanned int
	for ls.Scan() {
		var row Row

		// skip the first two lines of the file (header and blank line before data)
		if scanned++; scanned <= 2 {
			continue
		}

		l := strings.Replace(ls.Text(), "*", "", -1)
		if _, err := fmt.Fscan(strings.NewReader(l), &row.Day, &row.Max, &row.Min); err != nil {
			return nil, fmt.Errorf("unable to read values from line %d: %w", scanned, err)
		}

		rows = append(rows, row)
	}

	if err := ls.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan line %d: %w", scanned, err)
	}

	return rows, nil
}
