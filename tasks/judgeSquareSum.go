package tasks

import "math"

func judgeSquareSum(c int) bool {
	left := 0
	right := int(math.Sqrt(float64(c)))

	for left <= right {
		if left*left+right*right == c {
			return true
		}
		if left*left+right*right > c {
			right--
		} else {
			left--
		}
	}
	return false
}
