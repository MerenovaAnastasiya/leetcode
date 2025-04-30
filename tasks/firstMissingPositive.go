package tasks

func FirstMissingPositive(nums []int) int {
	i := 0
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] >= len(nums) || nums[i] < 1 || nums[i] == nums[j] {
			i++
		} else {
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	for j := 0; j < len(nums); j++ {
		if nums[j] != j+1 {
			return j + 1
		}
	}
	return len(nums) + 1
}
