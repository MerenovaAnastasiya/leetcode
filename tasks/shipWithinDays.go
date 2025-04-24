package tasks

func shipWithinDays(weights []int, days int) int {
	left := weights[0]
	right := 0

	for i := 0; i < len(weights); i++ {
		right += weights[i]
		if left < weights[i] {
			left = weights[i]
		}
	}

	if days == 1 {
		return right
	}

	result := right

	for left <= right {
		mid := left + (right-left)/2
		if canShip(weights, mid, days) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}

func canShip(weights []int, capacity int, days int) bool {
	currentLoad, requiredDays := 0, 1
	for _, weight := range weights {
		if currentLoad+weight > capacity {
			currentLoad = 0
			requiredDays++
		}
		currentLoad += weight
	}
	return requiredDays <= days
}
