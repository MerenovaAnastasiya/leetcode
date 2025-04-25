package tasks

import "strings"

func ReverseWords(s string) string {
	arr := strings.Fields(s)
	left := 0
	right := len(arr) - 1
	for left <= right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return strings.Join(arr, " ")
}
