package tasks

func MySqrt(x int) int {
	if x == 0 {
		return 0
	}
	low := 1
	high := x
	result := 0
	for low <= high {
		mid := (low + high) / 2
		if mid <= x/mid {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}
