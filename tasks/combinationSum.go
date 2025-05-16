package tasks

func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int

	var backtrack func([]int, []int, int, int)
	backtrack = func(path []int, candidates []int, target int, start int) {
		if target == 0 {
			result = append(result, append([]int{}, path...))
			return
		}
		if target < 0 {
			return
		}
		for i := start; i < len(candidates); i++ {
			target -= candidates[i]
			path = append(path, candidates[i])
			backtrack(path, candidates, target, i)
			target += candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtrack([]int{}, candidates, target, 0)
	return result
}
