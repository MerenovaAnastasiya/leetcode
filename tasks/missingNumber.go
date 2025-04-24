package tasks

func MissingNumber(nums []int) int {
	// n * (n + 1) / 2
	n := len(nums)
	seriesSum := n * (n + 1) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return seriesSum - sum
}
