package tasks

func longestPalindrome(s string) string {
	var start int
	end := len(s) / 2
	if len(s)%2 == 0 {
		start = len(s)/2 - 1

	} else {
		start = len(s) / 2
	}

	result := ""
	currentLength := 0
	maxLength := -1
	for start > 0 && end < len(s) {
		if s[start] == s[end] {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
				result = s[start:end]
			}
			currentLength = 0
		}
		start--
		end++
	}
	return result
}
