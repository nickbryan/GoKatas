package datamunging

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		data string
		want Rows
	}{
		"returns nil if nothing to parse": {
			data: ``,
			want: nil,
		},
		"parses a single record": {
			data: ` Dy MxT   MnT

   1  88    59`,
			want: Rows{
				Row{
					Day: 1,
					Min: 88,
					Max: 59,
				},
			},
		},
		"parses multiple records": {
			data: `  Dy MxT   MnT

   1  88    59
   2  79    63
   3  77    55`,
			want: Rows{
				Row{
					Day: 1,
					Min: 88,
					Max: 59,
				},
				Row{
					Day: 2,
					Min: 79,
					Max: 63,
				},
				Row{
					Day: 3,
					Min: 77,
					Max: 55,
				},
			},
		},
		"parses records with extra data": {
			data: `  Dy MxT   MnT   AvT   HDDay  AvDP 1HrP TPcpn WxType PDir AvSp Dir MxS SkyC MxR MnR AvSLP

   1  88    59    74          53.8       0.00 F       280  9.6 270  17  1.6  93 23 1004.5
   2  79    63    71          46.5       0.00         330  8.7 340  23  3.3  70 28 1004.5`,
			want: Rows{
				Row{
					Day: 1,
					Min: 88,
					Max: 59,
				},
				Row{
					Day: 2,
					Min: 79,
					Max: 63,
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Parse(strings.NewReader(tc.data))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
