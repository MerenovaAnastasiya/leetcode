package tasks

func FindDisappearedNumbers(nums []int) []int {
	start := 0
	for start < len(nums) {
		j := nums[start] - 1
		if nums[start] == nums[j] {
			start++
		} else {
			nums[start], nums[j] = nums[j], nums[start]
		}
	}
	var result []int
	for i, num := range nums {
		if num != i+1 {
			result = append(result, i+1)
		}
	}
	return result
}
