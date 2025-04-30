package tasks

func FindUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	n := len(nums)
	minNum := nums[n-1]
	maxNum := nums[0]

	l, r := -1, -1

	for i := 1; i < len(nums); i++ {
		if nums[i] >= maxNum {
			maxNum = nums[i]
		} else {
			r = i
		}
	}

	if r == -1 {
		return 0
	}

	for i := n - 2; i >= 0; i-- {
		if nums[i] <= minNum {
			minNum = nums[i]
		} else {
			l = i
		}
	}
	return r - l + 1

}
