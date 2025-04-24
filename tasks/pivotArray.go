package tasks

func PivotArray(nums []int, pivot int) []int {

	n := len(nums)
	output := make([]int, n)
	left, right := 0, n-1

	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		if nums[i] < pivot {
			output[left] = nums[i]
			left++
		}
		if nums[j] > pivot {
			output[right] = nums[j]
			right--
		}
	}

	for left <= right {
		output[left] = pivot
		left++
	}

	return output
}
