package tasks

import "strings"

func WordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")
	if len(arr) != len(pattern) {
		return false
	}
	mapS := map[string]byte{}
	mapP := map[byte]string{}

	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		word := arr[i]
		if val, ok := mapS[word]; ok && val != char {
			return false
		}
		if val, ok := mapP[char]; ok && val != word {
			return false
		}
		mapS[word] = char
		mapP[char] = word
	}
	return true
}
