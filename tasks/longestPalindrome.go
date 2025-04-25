package tasks

func longestPalindrome(s string) string {
	maxSubstring := ""
	for l := 0; l < len(s); l++ {
		for r := l + 1; r < len(s)+1; r++ {
			if r-l+1 > len(maxSubstring) {
				if check(s[l:r]) {
					maxSubstring = s[l:r]
				}
			}
		}
	}
	return maxSubstring
}

func check(s string) bool {
	n := len(s)
	start, end := 0, n-1

	for start < n/2 {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
