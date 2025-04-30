package tasks

func FindKthPositive(arr []int, k int) int {
	missingNumbersCount := 0
	j, current := 0, 1
	for missingNumbersCount < k {
		if j >= len(arr) || arr[j] != current {
			missingNumbersCount++
			if missingNumbersCount == k {
				return current
			}
		} else {
			j++
		}
		current++
	}
	return current

}
