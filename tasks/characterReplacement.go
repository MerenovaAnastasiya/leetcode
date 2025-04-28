package tasks

func CharacterReplacement(s string, k int) int {
	// AABABBA k = 1

	l, r, maxLength, maxCurrentLength := 0, 0, 0, 0
	var charCount [26]int

	for r < len(s) {
		charCount[s[r]-'A']++
		maxCurrentLength = max(maxCurrentLength, charCount[s[r]-'A'])
		otherChars := r - l + 1 - maxCurrentLength
		if otherChars > k {
			charCount[s[l]-'A']--
			l++
			maxCurrentLength = 0
		} else {
			maxLength = max(maxLength, r-l+1)
		}
		r++
	}
	return maxLength
}
