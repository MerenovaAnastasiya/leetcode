package tasks

import "slices"

func NumRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	count := 1
	currentWeight := 0
	r := 0
	// 1 2 2 3
	for r < len(people) {
		currentWeight += people[r]
		if currentWeight > limit {
			currentWeight = people[r]
			count++
		}
		r++
	}
	return count
}
