package sum_of_square_numbers

import (
	"math"
)

// link https://leetcode.com/problems/subtree-of-another-tree/description/
func judgeSquareSum(c int) bool {
	mid := math.Sqrt(float64(c) / 2)

	for i := 0.0; i <= mid; i++ {
		j := math.Sqrt(float64(c) - math.Pow(i, 2))
		if math.Floor(j) == j {
			return true
		}
	}

	return false
}
