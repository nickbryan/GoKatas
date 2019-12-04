package datamunging

import (
	"reflect"
	"testing"
)

func TestRow_Spread(t *testing.T) {
	tests := []struct {
		row  Row
		want float64
	}{
		{row: Row{Min: 0, Max: 0}, want: 0},
		{row: Row{Min: 2, Max: 2}, want: 0},
		{row: Row{Min: 0, Max: 1}, want: 1},
		{row: Row{Min: 0, Max: 2}, want: 2},
		{row: Row{Min: 1, Max: 10}, want: 9},
		{row: Row{Min: 100, Max: 120}, want: 20},
	}

	for _, tc := range tests {
		got := tc.row.Spread()
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("expected: %f, got: %f", tc.want, got)
		}
	}
}

func TestRows_MinSpread(t *testing.T) {
	tests := map[string]struct {
		rows Rows
		want *Row
	}{
		"returns nil if no records in set": {
			rows: Rows{},
			want: nil,
		},
		"returns the only record if set has single record": {
			rows: Rows{Row{Day: "1", Min: 1, Max: 6}},
			want: &Row{Day: "1", Min: 1, Max: 6},
		},
		"returns the record with minimum spread when set has multiple records": {
			rows: Rows{Row{Day: "1", Min: 1, Max: 16}, Row{Day: "2", Min: 1, Max: 3}},
			want: &Row{Day: "2", Min: 1, Max: 3},
		},
		"returns the most recent record if set has multiple records with matching min spread": {
			rows: Rows{Row{Day: "1", Min: 1, Max: 2}, Row{Day: "2", Min: 1, Max: 2}, Row{Day: "3", Min: 1, Max: 2}},
			want: &Row{Day: "3", Min: 1, Max: 2},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.rows.MinSpread()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
