package tasks

import "slices"

func Permute(nums []int) [][]int {
	var res [][]int
	var backtrack func([]int)
	backtrack = func(path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if slices.Contains(path, nums[i]) {
				continue
			}
			path = append(path, nums[i])
			backtrack(path)
			path = path[:len(path)-1]
		}

	}
	backtrack([]int{})
	return res

}
