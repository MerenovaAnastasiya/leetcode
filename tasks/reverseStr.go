package tasks

func ReverseStr(s string, k int) string {
	runes := []rune(s)

	for i := 0; i < len(runes); i += k * 2 {
		l := i
		r := i + k - 1
		for l <= r {
			if r >= len(runes) {
				r = len(runes) - 1
			}
			runes[l], runes[r] = runes[r], runes[l]
			l++
			r--
		}
	}
	return string(runes)
}
