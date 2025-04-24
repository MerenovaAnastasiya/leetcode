package tasks

func countGoodSubstrings(s string) int {
	m := make(map[byte]int)
	result, start, end := 0, 0, 2

	if len(s) < 3 {
		return result
	}

	for i := 0; i < 3; i++ {
		if value, ok := m[s[i]]; ok {
			m[s[i]] = value + 1
		} else {
			m[s[i]] = 1
		}
	}
	if len(m) == 3 {
		result++
	}

	for end < len(s)-1 {
		m[s[start]]--
		if m[s[start]] == 0 {
			delete(m, s[start])
		}
		start++
		end++
		if value, ok := m[s[end]]; ok {
			m[s[end]] = value + 1
		} else {
			m[s[end]] = 1
		}
		if len(m) == 3 {
			result++
		}
	}
	return result

}
