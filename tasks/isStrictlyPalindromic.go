package tasks

import "strconv"

func isStrictlyPalindromic(n int) bool {

	for i := 2; i <= n-2; i++ {
		currentString := toString(n, i)
		if !isPalindrome(currentString) {
			return false
		}
	}
	return true

}

func isPalindrome(str string) bool {
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func toString(n int, base int) string {
	result := ""
	for ; n > 0; n /= base {
		x := n % base
		result += strconv.Itoa(x)
	}
	return result
}
