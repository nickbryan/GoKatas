package karatechop

import (
	"math"
)

// NotFound will be returned when the needle can not be found within the given haystack.
const NotFound = -1

// BinarySearch is the function signature confirming to the kata specification.
type BinarySearch func(needle int, haystack []int) int

// IterativeBinarySearch uses a simple for loop and offsets to traverse the haystack when looking for the given needle.
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

// RecursiveBinarySearch uses slice splitting and recursion to find the needle in the given haystack. When the needle
// is in the upper half of the haystack, we must add the starting index of the upper half (mid point + 1) to the return
// value in order to return the correct index.
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

func TailRecursiveBinarySearch(needle int, haystack []int) int {
	// Define tailRecursiveBinarySearch upfront so we can call it recursively.
	// We could define this closure as a function elsewhere but this allows us to preserve encapsulation.
	var tailRecursiveBinarySearch func(needle int, haystack []int, min, max int) int

	tailRecursiveBinarySearch = func(needle int, haystack []int, min, max int) int {
		if max < min {
			return NotFound
		}

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

		return tailRecursiveBinarySearch(needle, haystack, min, max)
	}

	return tailRecursiveBinarySearch(needle, haystack, 0, len(haystack) - 1)
}
