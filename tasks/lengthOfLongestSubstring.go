package tasks

func lengthOfLongestSubstring(s string) int {
	l := 0
	longest := 0
	m := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		for m[s[i]] == true {
			m[s[l]] = false
			l++
		}
		longest = max(longest, i-l+1)
		m[s[i]] = true
	}
	return longest
}
