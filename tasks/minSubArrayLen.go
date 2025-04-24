package tasks

func LongestSubarray(nums []int) int {
	zeroes := 0
	maxLength := 0
	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			zeroes++
		}
		if zeroes > 1 {
			if nums[l] == 0 {
				zeroes--
			}
			l++

		}
		maxLength = max(maxLength, r-l+1)
	}
	return maxLength - 1
}
