package karatechop

import (
	"math"
)

func IterativeBinarySearch(needle int, haystack []int) int {
	min := 0
	max := len(haystack) - 1

	for max >= min {
		i := int(math.Floor(float64(min+max) / 2))

		if haystack[i] == needle {
			return i
		}

		if haystack[i] < needle {
			min = i + 1
		} else {
			max = i - 1
		}
	}

	return -1
}
