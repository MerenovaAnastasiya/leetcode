package tasks

func isHappy(n int) bool {
	if n == 1 || n == 7 {
		return true
	} else if n < 10 {
		return false
	} else {
		sum := 0
		for n > 0 {
			temp := n % 10
			sum += temp * temp
			n /= 10
		}
		return isHappy(sum)
	}
}
