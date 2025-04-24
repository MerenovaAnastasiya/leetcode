package tasks

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		element := nums[mid]
		if element == target {
			return mid
		} else if target < element {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
