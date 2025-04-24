package tasks

func longestNiceSubstring(s string) string {
	n := len(s)
	longest := ""

	for l := 0; l < n; l++ {
		lower := make([]bool, 26)
		upper := make([]bool, 26)

		for r := l; r < n; r++ {
			if s[r] >= 'a' && s[r] <= 'z' {
				lower[s[r]-'a'] = true
			} else {
				upper[s[r]-'A'] = true
			}

			niceString := true

			for i := 0; i < 26; i++ {
				if lower[i] && !upper[i] || !lower[i] && upper[i] {
					niceString = false
					break
				}
			}
			if niceString && r-l+1 > len(longest) {
				longest = s[l : r+1]
			}
		}

	}
	return longest

}
