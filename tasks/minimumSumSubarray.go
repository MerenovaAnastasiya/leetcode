package tasks

import "math"

func MinimumSumSubarray(nums []int, l int, r int) int {
	minSum := math.MaxInt32
	for i := l; i <= r; i++ {
		currentSum := 0

		for j := 0; j < i; j++ {
			currentSum += nums[j]
		}
		if currentSum > 0 {
			minSum = min(minSum, currentSum)
		}

		low, high := 0, i

		for high < len(nums) {
			currentSum -= nums[low]
			currentSum += nums[high]
			if currentSum > 0 {
				minSum = min(minSum, currentSum)
			}
			low++
			high++
		}
	}
	if minSum == math.MaxInt32 {
		return -1
	}
	return minSum
}
