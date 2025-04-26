package tasks

func CountSubstrings(s string) int {
	count := 0
	for l := 0; l < len(s); l++ {
		for r := l; r < len(s); r++ {
			if checkPalindrome(s[l : r+1]) {
				count++
			}
		}
	}
	return count
}

func checkPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
