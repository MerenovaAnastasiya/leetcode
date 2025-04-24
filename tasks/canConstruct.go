package tasks

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	var counts [26]int
	for _, r := range magazine {
		counts[r-'a']++
	}
	for _, r := range ransomNote {
		if counts[r-'a'] == 0 {
			return false
		}
		counts[r-'a']--
	}
	return true
}
