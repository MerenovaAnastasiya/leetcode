package tasks

func ReverseVowels(s string) string {
	m := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	left := 0
	right := len(s) - 1
	b := []byte(s)
	for left < right {
		if m[b[left]] && m[b[right]] {
			temp := s[left]
			b[left] = b[right]
			b[right] = temp
			left++
			right--
		} else if !m[b[left]] {
			left++
		} else {
			right--
		}
	}
	return string(b)
}
