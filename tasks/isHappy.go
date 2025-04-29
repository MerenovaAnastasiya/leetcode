package tasks

func isHappy(n int) bool {
	slow, fast := n, n
	for fast != slow {
		fast = findSquare(findSquare(fast))
		slow = findSquare(slow)
	}
	return fast == 1
}

func findSquare(n int) int {
	sum := 0
	for n > 0 {
		temp := n % 10
		sum += temp * temp
		n /= 10
	}
	return sum
}
