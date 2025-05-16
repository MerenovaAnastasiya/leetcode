package tasks

func getRow(rowIndex int) []int {
	prev := []int{1}
	for i := 1; i <= rowIndex; i++ {
		nums := make([]int, i+1)
		nums[0] = 1
		for j := 0; j < i-1; j++ {
			nums = append(nums, prev[j]+prev[j+1])
		}
		nums = append(nums, 1)
		prev, nums = nums, prev
	}
	return prev
}
