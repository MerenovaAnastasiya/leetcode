package tasks

func MaximumLengthSubstring(s string) int {
	m := make(map[byte]int)
	maxLength, l := 0, 0

	for r := 0; r < len(s); r++ {
		m[s[r]]++
		for m[s[r]] > 2 {
			m[s[l]]--
			l++
		}
		maxLength = max(maxLength, r-l+1)
	}
	return maxLength

}
