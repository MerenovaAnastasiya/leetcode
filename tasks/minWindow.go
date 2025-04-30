package tasks

import "math"

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	var alphabet [128]int

	for i := 0; i < len(t); i++ {
		alphabet[t[i]-'a']++
	}
	l, r, minStart, count, minLen := 0, 0, 0, len(t), math.MaxInt32

	for r < len(s) {
		if alphabet[s[r]-'a'] > 0 {
			count--
		}
		alphabet[s[r]-'a']--

		for count == 0 {
			if minLen > r-l+1 {
				minLen = r - l + 1
				minStart = l
			}
			alphabet[s[l]-'a']++
			if alphabet[s[l]-'a'] > 0 {
				count++
			}
			l++
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[minStart : minStart+minLen]
}
