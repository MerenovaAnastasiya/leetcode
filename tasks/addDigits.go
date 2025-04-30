package tasks

func AddDigits(num int) int {
	sum := 0
	for {
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		num, sum = sum, 0
		if num < 10 {
			break
		}

	}

	return sum
}
