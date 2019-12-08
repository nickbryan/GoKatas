package datamunging

import (
	"fmt"
	"io"
)

// MinSpreadCalculator returns the Row with the minimum spread.
type MinSpreadCalculator interface {
	MinSpread() Row
}

// WriteMinSpread uses the given MinSpreadCalculator to write the Row with the minimum spread to the given io.Writer.
// The Id and MinSpread are written out with a newline.
func WriteMinSpread(c MinSpreadCalculator, w io.Writer) error {
	row := c.MinSpread()
	_, err := w.Write([]byte(fmt.Sprintf("%s -- Spread: %f\n", row.Id, row.Spread())))
	return err
}
