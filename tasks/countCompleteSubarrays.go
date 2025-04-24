package tasks

func countCompleteSubarrays(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	count := len(m)
	return b(nums, count) - b(nums, count-1)
}

func b(nums []int, x int) int {
	l, count := 0, 0
	n := len(nums)
	m := make(map[int]int)
	for r := 0; r < n; r++ {
		m[nums[r]]++
		for len(m) > x {
			if m[nums[l]] == 1 {
				delete(m, nums[l])
			} else {
				m[nums[l]]--
			}
			l++
		}
		count += r - l + 1
	}
	return count
}
