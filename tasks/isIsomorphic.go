package tasks

func isIsomorphic(s string, t string) bool {
	m1 := map[byte]byte{}
	// s -> t

	m2 := map[byte]byte{}

	// t - > s
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		ch1 := s[i]
		ch2 := t[i]

		if val, ok := m1[ch1]; ok && val != ch2 {
			return false
		}
		if val, ok := m2[ch2]; ok && val != ch1 {
			return false
		}
		m1[ch1] = ch2
		m2[ch2] = ch1
	}
	return true
}
