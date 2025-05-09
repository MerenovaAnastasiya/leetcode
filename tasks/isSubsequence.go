package tasks

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	k := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[k] {
			k++
			if k == len(s) {
				return true
			}
		}
	}
	return false
}
