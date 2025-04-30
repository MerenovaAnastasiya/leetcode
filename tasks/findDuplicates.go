package tasks

func FindDuplicates(nums []int) []int {
	i := 0
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] == nums[j] {
			i++
		} else {
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	var result []int
	for j := 0; j < len(nums); j++ {
		if nums[j] != j+1 {
			result = append(result, nums[j])
		}
	}
	return result
}
