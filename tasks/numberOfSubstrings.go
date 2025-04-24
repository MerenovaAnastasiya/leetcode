package tasks

func NumberOfSubstrings(s string) int {
	count := 0
	m := make([]int, 3)
	left, right := 0, 0
	for right < len(s) {
		m[s[right]-'a']++

		for allChars(m) {
			count += len(s) - right
			m[s[left]-'a']--
			left++
		}
		right++
	}
	return count
}

func allChars(m []int) bool {
	flg := true
	for i := 0; i < len(m); i++ {
		if m[i] <= 0 {
			flg = false
			break
		}
	}
	return flg
}
