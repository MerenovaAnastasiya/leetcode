package tasks

func generateParenthesis(n int) []string {
	var queue []Str
	var result []string
	queue = append(queue, Str{"", 0, 0})
	for len(queue) > 0 {
		current := queue[0]
		if current.lCount == n && current.rCount == n {
			result = append(result, current.s)
		} else {
			if current.lCount < n {
				queue = append(queue, Str{current.s + "(", current.lCount + 1, current.rCount})
			}
			if current.rCount < current.lCount {
				queue = append(queue, Str{current.s + ")", current.lCount, current.rCount + 1})
			}
		}
		queue = queue[1:]
	}
	return result
}

type Str struct {
	s      string
	lCount int
	rCount int
}
