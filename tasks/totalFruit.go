package tasks

func totalFruit(fruits []int) int {
	l, maxLength, maxCount := 0, 0, 2
	m := make(map[int]int)

	// 1,0,1,4,1,4,1,2,3
	for r := 0; r < len(fruits); r++ {
		m[fruits[r]]++
		for len(m) > maxCount {
			m[fruits[l]]--
			if m[fruits[l]] == 0 {
				delete(m, fruits[l])
			}
			l++
		}
		maxLength = max(r-l+1, maxLength)
	}
	return maxLength
}
