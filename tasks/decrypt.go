package tasks

func Decrypt(code []int, k int) []int {
	n := len(code)
	result := make([]int, n)

	if k == 0 {
		return result
	}

	start, end := 1, k

	if k < 0 {
		start, end = n-abs(k), n-1
	}

	sum := 0
	for i := start; i <= end; i++ {
		sum += code[i]
	}

	for i := 0; i < n; i++ {
		result[i] = sum
		sum -= code[start%n]
		sum += code[(end+1)%n]
		start++
		end++
	}

	return result
}

func abs(k int) int {
	if k > 0 {
		return k
	} else {
		return -k
	}
}
