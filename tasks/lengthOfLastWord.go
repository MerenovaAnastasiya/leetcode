package tasks

func lengthOfLastWord(s string) int {
	end := len(s) - 1
	hasWord := false
	for i := len(s) - 1; i >= 0; i-- {
		if !hasWord && s[i] != ' ' {
			end = i
			hasWord = true
		}
		if hasWord && (s[i] == ' ') {
			return end - i
		}
	}
	return end + 1
}
