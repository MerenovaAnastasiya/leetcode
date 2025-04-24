package tasks

import "fmt"

func NumSubarraysWithSum(nums []int, goal int) int {
	return atMost(nums, goal) - atMost(nums, goal-1)
}

func atMost(nums []int, goal int) int {
	// считаем количество подмассивов, сумма которых <=2

	left, count, sum := 0, 0, 0
	n := len(nums)
	for right := 0; right < n; right++ {
		sum += nums[right]
		for sum > goal && left <= right {
			sum -= nums[left]
			left++
		}
		count += right - left + 1
		fmt.Printf("%v", nums[left:right+1])
	}
	return count
}
