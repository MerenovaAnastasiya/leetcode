package tasks

func generate(numRows int) [][]int {
	dp := [][]int{{1}}

	for i := 1; i < numRows; i++ {
		nums := []int{1}
		for j := 0; j < i-1; j++ {
			nums = append(nums, dp[i-1][j]+dp[i-1][j+1])
		}
		nums = append(nums, 1)
		dp = append(dp, nums)
	}
	return dp
}
