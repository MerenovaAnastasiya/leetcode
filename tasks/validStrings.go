package tasks

func ValidStrings(n int) []string {
	var result []string

	var backtrack func(int, string, string)
	backtrack = func(index int, str string, str2 string) {
		str = str + str2
		if index == n {
			result = append(result, str)
			return
		}
		if str2 == "0" {
			backtrack(index+1, str, "1")
		} else {
			backtrack(index+1, str, "1")
			backtrack(index+1, str, "0")
		}
		if len(str) != 0 {
			str = str[:len(str)-1]
		}

	}
	backtrack(0, "", "")

	return result
}
