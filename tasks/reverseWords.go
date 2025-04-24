package tasks

func ReverseWords(s string) string {
	start := 0
	result := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			end := i - 1
			for start < i {
				current := s[start]
				result[start] = s[end]
				result[end] = current
				start++
				end--
			}
			start = i + 1
		}
	}
	end := len(result) - 1
	for start < len(s) {
		current := s[start]
		result[start] = s[end]
		result[end] = current
		start++
		end--
	}
	return string(result)
}
