package tasks

func NumberOfSubarrays(nums []int, k int) int {

	result := 0
	left, right, count, temp := 0, 0, 0, 0

	for right < len(nums) {
		if nums[right]%2 == 1 {
			count++
			temp = 0
		}
		for count == k {
			temp++
			if nums[left]%2 == 1 {
				count--
			}
			left++
		}
		result += temp
		right++
	}
	return result
}
