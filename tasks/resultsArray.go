package tasks

func ResultsArray(nums []int, k int) []int {
	result := make([]int, len(nums)-k+1)

	if len(nums) == 1 || k == 1 {
		return nums
	}
	count := 1
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			count++
		} else {
			count = 1
		}

		if count >= k {
			result[j] = nums[i]
		} else if i < k-1 {
			continue
		} else {
			result[j] = -1
		}
		j++
	}

	return result

}
