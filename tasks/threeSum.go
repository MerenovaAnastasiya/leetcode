package tasks

import (
	"math"
	"slices"
)

func ThreeSum(nums []int) [][]int {
	slices.Sort(nums)
	sum := math.MaxInt32
	var result [][]int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum = nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			}
		}
	}
	return result

}
