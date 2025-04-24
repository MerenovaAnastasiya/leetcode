package tasks

import "slices"

func CountPairs(nums []int, target int) int {
	result := 0
	if len(nums) < 2 {
		return result
	}

	slices.Sort(nums)
	l, r := 0, len(nums)-1

	for l < r {
		sum := nums[l] + nums[r]
		if sum < target {
			result += r - l
			l++
		} else {
			r--
		}
	}
	return result

}
