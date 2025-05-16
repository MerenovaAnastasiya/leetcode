package tasks

import (
	"strconv"
)

func isBalanced(num string) bool {
	evenSum, oddSum := 0, 0
	evenPinter, oddPinter := 1, 0

	for oddPinter < len(num) {
		digit, err := strconv.Atoi(string(num[oddPinter]))
		if err != nil {
			continue
		}
		oddSum += digit
		oddPinter += 2
	}

	for evenPinter < len(num) {
		digit, err := strconv.Atoi(string(num[evenPinter]))
		if err != nil {
			continue
		}
		evenSum += digit
		evenPinter += 2
	}

	return oddSum == evenSum
}
