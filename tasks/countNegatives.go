package tasks

func countNegatives(grid [][]int) int {
	count, l := 0, 0
	for i := len(grid) - 1; i >= 0; i-- {
		l = bs(grid[i], l, len(grid[i])-1)
		if l == len(grid[i]) {
			break
		}
		count += len(grid[i]) - l
	}
	return count
}

func bs(arr []int, left int, right int) int {
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
