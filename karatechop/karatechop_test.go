package karatechop

import (
	"testing"
)

func TestIterativeBinarySearch(t *testing.T) {
	runTests(t, IterativeBinarySearch)
}

func BenchmarkIterativeBinarySearch(b *testing.B) {
	runBenchmarks(b, IterativeBinarySearch)
}

func TestRecursiveBinarySearch(t *testing.T) {
	runTests(t, RecursiveBinarySearch)
}

func BenchmarkRecursiveBinarySearch(b *testing.B) {
	runBenchmarks(b, RecursiveBinarySearch)
}

func TestTailRecursiveBinarySearch(t *testing.T) {
	runTests(t, TailRecursiveBinarySearch)
}

func BenchmarkTailRecursiveBinarySearch(b *testing.B) {
	runBenchmarks(b, TailRecursiveBinarySearch)
}

func TestParallelIterativeBinarySearch(t *testing.T) {
	runTests(t, ParallelIterativeBinarySearch)
}

func BenchmarkParallelIterativeBinarySearch(b *testing.B) {
	runBenchmarks(b, ParallelIterativeBinarySearch)
}

var tests = map[string]struct {
	input, want int
	set         []int
}{
	"target not in empty set":          {input: 3, set: []int{}, want: NotFound},
	"target not in single element set": {input: 3, set: []int{1}, want: NotFound},
	"target in single element set":     {input: 1, set: []int{1}, want: 0},

	"target at start of 3 element set":  {input: 1, set: []int{1, 3, 5}, want: 0},
	"target in middle of 3 element set": {input: 3, set: []int{1, 3, 5}, want: 1},
	"target at end of set":              {input: 5, set: []int{1, 3, 5}, want: 2},
	"target not in 3 element set 1":     {input: 0, set: []int{1, 3, 5}, want: NotFound},
	"target not in 3 element set 2":     {input: 2, set: []int{1, 3, 5}, want: NotFound},
	"target not in 3 element set 3":     {input: 4, set: []int{1, 3, 5}, want: NotFound},
	"target not in 3 element set 4":     {input: 6, set: []int{1, 3, 5}, want: NotFound},

	"target at start of 4 element set":     {input: 1, set: []int{1, 3, 5, 7}, want: 0},
	"target at 2nd index of 4 element set": {input: 3, set: []int{1, 3, 5, 7}, want: 1},
	"target at 3rd index of 4 element set": {input: 5, set: []int{1, 3, 5, 7}, want: 2},
	"target at end of 4 element set":       {input: 7, set: []int{1, 3, 5, 7}, want: 3},
	"target not in 4 element set 1":        {input: 0, set: []int{1, 3, 5, 7}, want: NotFound},
	"target not in 4 element set 2":        {input: 2, set: []int{1, 3, 5, 7}, want: NotFound},
	"target not in 4 element set 3":        {input: 4, set: []int{1, 3, 5, 7}, want: NotFound},
	"target not in 4 element set 4":        {input: 6, set: []int{1, 3, 5, 7}, want: NotFound},
	"target not in 4 element set 5":        {input: 8, set: []int{1, 3, 5, 7}, want: NotFound},
}

func runTests(t *testing.T, algorithm BinarySearch) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := algorithm(tc.input, tc.set)
			if got != tc.want {
				t.Errorf("expected: %d, got: %d", tc.want, got)
			}
		})
	}
}

var (
	benchmarks = map[string]struct {
		len, target int
	}{
		"first index large":  {len: 1e6, target: 1},
		"last index large":   {len: 1e6, target: 1e6},
		"middle index large": {len: 1e6, target: 0.5e6},
		"upper half large":   {len: 1e6, target: 0.8e6},
		"lower half large":   {len: 1e6, target: 0.2e6},

		"first index medium":  {len: 1e5, target: 1},
		"last index medium":   {len: 1e5, target: 1e5},
		"middle index medium": {len: 1e5, target: 0.5e5},
		"upper half medium":   {len: 1e5, target: 0.8e5},
		"lower half medium":   {len: 1e5, target: 0.2e5},

		"first index small":  {len: 1e4, target: 1},
		"last index small":   {len: 1e4, target: 1e5},
		"middle index small": {len: 1e4, target: 0.5e4},
		"upper half small":   {len: 1e4, target: 0.8e4},
		"lower half small":   {len: 1e4, target: 0.2e4},
	}

	// result is used to prevent the compiler from eliminating any Benchmarks during optimisations by storing th result
	// to a package level variable.
	result int
)

func runBenchmarks(b *testing.B, algorithm BinarySearch) {
	for name, bm := range benchmarks {
		b.Run(name, func(b *testing.B) {
			s := makeSlice(bm.len)
			var r int
			for n := 0; n < b.N; n++ {
				// store result locally to prevent compiler eliminating function call.
				r = algorithm(bm.target, s)
			}
			result = r
		})
	}
}

func makeSlice(len int) []int {
	s := make([]int, len)
	for i := range s {
		s[i] = i + 1
	}
	return s
}
