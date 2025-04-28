package tasks

import (
	"math"
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)
	minDiff := math.MaxInt32
	result := 0

	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < minDiff {
				minDiff = abs(sum - target)
				result = sum
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return result
}
