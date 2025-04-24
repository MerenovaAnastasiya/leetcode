package tasks

func PartitionLabels(s string) []int {
	last := make([]int, 26)

	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	start, end := 0, 0
	var res []int
	for i := 0; i < len(s); i++ {
		end = max(end, last[s[i]-'a'])
		if end == i {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}
