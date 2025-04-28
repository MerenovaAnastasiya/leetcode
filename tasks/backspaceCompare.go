package tasks

func BackspaceCompare(s string, t string) bool {
	// s = "ab#c", t = "ad#c"
	str1 := []rune(s)
	str2 := []rune(t)
	res1 := createStr(str1)
	res2 := createStr(str2)

	if res1 != res2 {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}

func createStr(str []rune) int {
	k := 0
	for i := 0; i < len(str); i++ {
		if str[i] != '#' {
			str[k] = str[i]
			k++
		} else if k > 0 {
			k--
		}
	}
	return k
}
