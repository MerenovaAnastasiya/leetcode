package tasks

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	count := len(fruits)
	used := make([]bool, len(baskets))
	for _, fruitCount := range fruits {
		for j, basketSize := range baskets {
			if basketSize >= fruitCount {
				if !used[j] {
					used[j] = true
					count--
					break
				}
			}
		}
	}
	return count
}
