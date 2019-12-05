package datamunging

import (
	"fmt"
	"io"
)

// MinSpreader returns the Row with the minimum spread.
type MinSpreader interface {
	MinSpread() *Row
}

// WriteMinSpread uses the given MinSpreader to write the Row with the minimum spread to the given io.Writer.
// The Day and minimum spread are written out with a newline.
func WriteMinSpread(c MinSpreader, w io.Writer) error {
	row := c.MinSpread()
	_, err := w.Write([]byte(fmt.Sprintf("Day: %s, Min Spread: %f\n", row.Day, row.Spread())))
	return err
}
