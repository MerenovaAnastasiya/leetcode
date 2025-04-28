package tasks

func LongestOnes(nums []int, k int) int {
	l, r, length, count := 0, 0, 0, 0
	for r < len(nums) {
		if nums[r] == 0 {
			count++
		}
		for count > k {
			if nums[l] == 0 {
				count--
			}
			l++
		}
		length = max(length, r-l+1)
		r++
	}
	return length
}
