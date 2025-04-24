package tasks

func StrStr(haystack string, needle string) int {
	k, l := 0, 0
	for l < len(haystack) {
		if needle[k] == haystack[l] {
			k++
			l++
		} else {
			l = l - k + 1
			k = 0
		}
		if k == len(needle) {
			return l - k
		}
	}
	return -1

}
