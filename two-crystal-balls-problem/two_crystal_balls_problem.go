package twocrystalballsproblem

import (
	"math"
)

func twoCrystalBalls(haystack []bool) int {
	if len(haystack) == 0 {
		return -1
	}

	jumpAmount := int(math.Sqrt(float64(len(haystack))))
	index := jumpAmount
	for i := jumpAmount; i < len(haystack); i += jumpAmount {
		index = i
		if haystack[i] == true {
			break
		}
	}

	index -= jumpAmount

	for i, j := index, 0; j < jumpAmount && i <= len(haystack); i, j = i+1, j+1 {
		if haystack[i] == true {
			return i
		}
	}

	return -1
}
