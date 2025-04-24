package tasks

func rotate(nums []int, k int) {
	k %= len(nums)
	// переворачиваем весь массив
	reverseArray(nums, 0, len(nums)-1)
	// переворачиваем часть от 0 до k - 1
	reverseArray(nums, 0, k-1)
	// переворачиваем вторую часть(от k до n)
	reverseArray(nums, k, len(nums)-1)
}

func reverseArray(nums []int, start int, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
