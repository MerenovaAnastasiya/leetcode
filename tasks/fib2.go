package tasks

func fib2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	a, b := 0, 1

	for i := 2; i <= n; i++ {
		temp := a + b
		a = b
		b = temp

	}
	return b
}
