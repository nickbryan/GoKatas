package datamunging

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Row struct {
	Day      uint
	Min, Max float64
}

type Rows []Row

func Parse(r io.Reader) Rows {
	ls := bufio.NewScanner(r)
	ls.Split(bufio.ScanLines)
	ls.Scan()
	ls.Scan()

	var rows Rows
	for ls.Scan() {
		ws := bufio.NewScanner(strings.NewReader(ls.Text()))
		ws.Split(bufio.ScanWords)

		var records [3]string
		for i := 0; i < 3; i++ {
			ws.Scan()
			records[i] = ws.Text()
		}

		d, err := strconv.ParseUint(records[0], 10, 32)
		if err != nil {
			return nil
		}

		mn, err := strconv.ParseFloat(records[1], 64)
		if err != nil {
			return nil
		}

		mx, err := strconv.ParseFloat(records[2], 64)
		if err != nil {
			return nil
		}

		rows = append(rows, Row{
			Day: uint(d),
			Min: mn,
			Max: mx,
		})
	}
	//if err := scanner.Err(); err != nil {
	//		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	//	}

	return rows
}
