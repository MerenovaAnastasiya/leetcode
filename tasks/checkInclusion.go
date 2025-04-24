package tasks

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	str1, str2 := make([]int, 26), make([]int, 26)

	for _, char := range s1 {
		str1[char-'a']++
	}

	for i, char := range s2 {
		str2[char-'a']++
		if i >= len(s1) {
			str2[s2[i-len(s1)]-'a']--
		}
		if isPermutation(str1, str2) {
			return true
		}
	}
	return false
}

func isPermutation(str1 []int, str2 []int) bool {
	for i, v := range str1 {
		if v != str2[i] {
			return false
		}
	}
	return true
}
