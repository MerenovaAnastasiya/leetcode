package tasks

import "slices"

func FindContentChildren(g []int, s []int) int {
	// g - children
	// s - cookies

	slices.Sort(g)
	slices.Sort(s)

	count := 0
	k, l := 0, 0

	if len(s) == 0 {
		return 0
	}

	for l < len(g) {
		if s[k] >= g[l] {
			count++
			l++
		}
		k++
		if k == len(s) {
			break
		}
	}
	return count
}
