package tasks

func maxProfit(prices []int) int {
	maxProfit := 0
	minValue := prices[0]
	for i := 1; i < len(prices); i++ {
		current := prices[i]
		if current < minValue {
			minValue = current
		}
		if current-minValue > maxProfit {
			maxProfit = current - minValue
		}
	}
	return maxProfit
}
