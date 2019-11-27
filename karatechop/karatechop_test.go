package karatechop

import "testing"

var tests = map[string]struct {
	input int
	set   []int
	want  int
}{
	"target not in empty set":          {input: 3, set: []int{}, want: -1},
	"target not in single element set": {input: 3, set: []int{1}, want: -1},
	"target in single element set":     {input: 1, set: []int{1}, want: 0},

	"target at start of 3 element set":  {input: 1, set: []int{1, 3, 5}, want: 0},
	"target in middle of 3 element set": {input: 3, set: []int{1, 3, 5}, want: 1},
	"target at end of set":              {input: 5, set: []int{1, 3, 5}, want: 2},
	"target not in 3 element set 1":     {input: 0, set: []int{1, 3, 5}, want: -1},
	"target not in 3 element set 2":     {input: 2, set: []int{1, 3, 5}, want: -1},
	"target not in 3 element set 3":     {input: 4, set: []int{1, 3, 5}, want: -1},
	"target not in 3 element set 4":     {input: 6, set: []int{1, 3, 5}, want: -1},

	"target at start of 4 element set":     {input: 1, set: []int{1, 3, 5, 7}, want: 0},
	"target at 2nd index of 4 element set": {input: 3, set: []int{1, 3, 5, 7}, want: 1},
	"target at 3rd index of 4 element set": {input: 5, set: []int{1, 3, 5, 7}, want: 2},
	"target at end of 4 element set":       {input: 7, set: []int{1, 3, 5, 7}, want: 3},
	"target not in 4 element set 1":        {input: 0, set: []int{1, 3, 5, 7}, want: -1},
	"target not in 4 element set 2":        {input: 2, set: []int{1, 3, 5, 7}, want: -1},
	"target not in 4 element set 3":        {input: 4, set: []int{1, 3, 5, 7}, want: -1},
	"target not in 4 element set 4":        {input: 6, set: []int{1, 3, 5, 7}, want: -1},
	"target not in 4 element set 5":        {input: 8, set: []int{1, 3, 5, 7}, want: -1},
}

func TestIterativeBinarySearch(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := IterativeBinarySearch(tc.input, tc.set)
			if got != tc.want {
				t.Errorf("expected: %d, got: %d", tc.want, got)
			}
		})
	}
}
