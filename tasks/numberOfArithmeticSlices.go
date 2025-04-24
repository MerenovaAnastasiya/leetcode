package tasks

func NumberOfArithmeticSlices(nums []int) int {
	n := len(nums)

	// кол-во подмассивов, в которых разница между соседними эл-ми одинаковая

	if n < 3 {
		return 0
	}
	// 1,2,3,4,5,6
	right, left, result := 2, 0, 0
	k := 3
	diff := nums[1] - nums[0]

	for right < n {
		currentDiff := nums[right] - nums[right-1]
		for currentDiff != diff && left < right-1 {
			left++
			diff = nums[left+1] - nums[left]
		}
		if right-left+1 >= k {
			result += right - left - 1
		}
		right++
	}

	return result
}
