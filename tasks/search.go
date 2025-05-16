package tasks

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		med := l + (r-l)/2
		if nums[med] == target {
			return med
		}
		if nums[med] < target {
			l = med + 1
		} else {
			r = med - 1
		}
	}
	return -1
}
