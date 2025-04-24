package tasks

func removeDuplicates(nums []int) int {
	k := 0
	current := nums[0]
	for i := 1; i < len(nums); i++ {
		if current != nums[i] {
			k++
			nums[k] = nums[i]
			current = nums[i]
		}
	}
	return k
}
