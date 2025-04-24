package tasks

func IsLongPressedName(name string, typed string) bool {
	l, k := 0, 0
	for k < len(typed) {
		if l < len(name) && typed[k] == name[l] {
			l++
			k++
		} else if k > 0 && typed[k] == typed[k-1] {
			k++
		} else {
			return false
		}
	}
	return l == len(name)
}
