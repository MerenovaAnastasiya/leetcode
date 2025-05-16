package tasks

func DivisorGame(n int) bool {
	var memo = map[int]bool{}
	var rec func(int) bool
	rec = func(x int) bool {
		if x == 1 {
			return false
		}
		if can, found := memo[n]; found {
			return can
		}
		for i := 1; i < x; i++ {
			if x%i == 0 && !rec(x-1) {
				memo[x] = true
				return true
			}
		}
		memo[x] = false
		return false
	}
	rec(n)
	return memo[n]
}
