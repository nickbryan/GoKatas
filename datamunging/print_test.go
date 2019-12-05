package datamunging

import (
	"bytes"
	"fmt"
	"testing"
)

type mockMinSpreadCalculator struct {
	row *Row
}

func (m mockMinSpreadCalculator) MinSpread() *Row {
	return m.row
}

func TestWriteMinSpread(t *testing.T) {
	row := &Row{
		Day: "1",
		Min: 5,
		Max: 10,
	}
	want := fmt.Sprintf("Day: %s, Min Spread: %f\n", row.Day, row.Spread())

	w := new(bytes.Buffer)
	if err := WriteMinSpread(mockMinSpreadCalculator{row: row}, w); err != nil {
		t.Fatalf("unable to WriteMinSpread: %v", err)
	}

	if w.String() != want {
		t.Errorf("expected: %s, got: %s", want, w.String())
	}
}
