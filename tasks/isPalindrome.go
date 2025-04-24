package tasks

import "strings"
import _ "unicode/utf8"

func IsPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	s = strings.ToLower(s)
	for left <= right {
		if (s[left] > 'z' || s[left] < 'a') && (s[left] < '0' || s[left] > '9') {
			left++
			continue
		}
		if (s[right] > 'z' || s[right] < 'a') && (s[right] < '0' || s[right] > '9') {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
