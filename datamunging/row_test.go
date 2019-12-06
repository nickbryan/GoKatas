package datamunging

import (
	"reflect"
	"strings"
	"testing"
)

func TestRow_Spread(t *testing.T) {
	tests := []struct {
		row  *Row
		want float64
	}{
		{row: &Row{A: 0, B: 0}, want: 0},
		{row: &Row{A: 2, B: 2}, want: 0},
		{row: &Row{A: 0, B: 1}, want: 1},
		{row: &Row{A: 0, B: 2}, want: 2},
		{row: &Row{A: 1, B: 10}, want: 9},
		{row: &Row{A: 100, B: 120}, want: 20},
	}

	for _, tc := range tests {
		got := tc.row.Spread()
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("expected: %f, got: %f", tc.want, got)
		}
	}
}

func TestRows_ReadFrom(t *testing.T) {
	tests := map[string]struct {
		data string
		want Rows
	}{
		"returns empty rows if nothing to parse": {
			data: ``,
			want: Rows{},
		},
		"parses a single record": {
			data: ` Dy MxT   MnT

   1  88    59`,
			want: Rows{
				&Row{
					Id: "1",
					A:  88,
					B:  59,
				},
			},
		},
		"parses non numeric identifiers": {
			data: ` Dy MxT   MnT

   mo  88    59`,
			want: Rows{
				&Row{
					Id: "mo",
					A:  88,
					B:  59,
				},
			},
		},
		"parses multiple records": {
			data: `  Dy MxT   MnT

   1  88    59
   2  79    63
   3  77    55`,
			want: Rows{
				&Row{
					Id: "1",
					A:  88,
					B:  59,
				},
				&Row{
					Id: "2",
					A:  79,
					B:  63,
				},
				&Row{
					Id: "3",
					A:  77,
					B:  55,
				},
			},
		},
		"parses records with extra data": {
			data: `  Dy MxT   MnT   AvT   HDDay  AvDP 1HrP TPcpn WxType PDir AvSp Dir MxS SkyC MxR MnR AvSLP

   1  88    59    74          53.8       0.00 F       280  9.6 270  17  1.6  93 23 1004.5
   2  79    63    71          46.5       0.00         330  8.7 340  23  3.3  70 28 1004.5`,
			want: Rows{
				&Row{
					Id: "1",
					A:  88,
					B:  59,
				},
				&Row{
					Id: "2",
					A:  79,
					B:  63,
				},
			},
		},
		"parses data in a different format": {
			data: `Team            P     W    L   D    F      A     Pts
Arsenal         38    26   9   3    79    36    87`,
			want: Rows{
				&Row{
					Id: "Arsenal",
					A:  79,
					B:  36,
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Rows{}
			err := got.ReadFrom(strings.NewReader(tc.data))
			if err != nil {
				t.Fatalf("an error occured on Parse: %v", err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestRows_MinSpread(t *testing.T) {
	tests := map[string]struct {
		rows Rows
		want *Row
	}{
		"returns zero value Row if no records in set": {
			rows: Rows{},
			want: &Row{},
		},
		"returns the only record if set has single record": {
			rows: Rows{&Row{Id: "1", A: 1, B: 6}},
			want: &Row{Id: "1", A: 1, B: 6},
		},
		"returns the record with minimum spread when set has multiple records": {
			rows: Rows{&Row{Id: "1", A: 1, B: 16}, &Row{Id: "2", A: 1, B: 3}},
			want: &Row{Id: "2", A: 1, B: 3},
		},
		"returns the most recent record if set has multiple records with matching min spread": {
			rows: Rows{&Row{Id: "1", A: 1, B: 2}, &Row{Id: "2", A: 1, B: 2}, &Row{Id: "3", A: 1, B: 2}},
			want: &Row{Id: "3", A: 1, B: 2},
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
