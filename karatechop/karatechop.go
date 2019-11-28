package karatechop

import (
	"math"
)

const NotFound = -1

func IterativeBinarySearch(needle int, haystack []int) int {
	min, max := 0, len(haystack)-1

	for max >= min {
		i := int(math.Floor(float64(min+max) / 2))
		v := haystack[i]

		switch {
		case v == needle:
			return i
		case v < needle:
			min = i + 1
		case v > needle:
			max = i - 1
		}
	}

	return NotFound
}

func RecursiveBinarySearch(needle int, haystack []int) int {
	min, max := 0, len(haystack)-1

	if max >= min {
		i := int(math.Floor(float64(min+max) / 2))
		v := haystack[i]

		switch {
		case v == needle:
			return i
		case v < needle:
			nextMin := i + 1
			r := RecursiveBinarySearch(needle, haystack[nextMin:])

			if r != NotFound {
				r += nextMin
			}

			return r
		case v > needle:
			return RecursiveBinarySearch(needle, haystack[:i])
		}
	}

	return NotFound
}
