package tasks

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		if count[t[i]-'a'] < 0 {
			return false
		}
		count[t[i]-'a']--
	}
	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			return false
		}
	}
	return true
}
