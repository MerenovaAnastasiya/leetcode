package tasks

import "slices"

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	// [1, 2, 2]
	var res [][]int

	var backtrack func(start int, path []int)

	backtrack = func(start int, path []int) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			if i > 0 && i < start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, []int{})
	return res

}
