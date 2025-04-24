package tasks

func equalSubstring(s string, t string, maxCost int) int {
	l := 0
	maxLength := -1
	currentCost := 0
	for r := 0; r < len(s); r++ {
		currentCost += abs(int(s[r]) - int(t[r]))
		for currentCost > maxCost {
			currentCost -= abs(int(s[l]) - int(t[l]))
			l++
		}
		maxLength = max(maxLength, r-l+1)
	}
	return maxLength

}
