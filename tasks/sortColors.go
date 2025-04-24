package tasks

func SortColors(nums []int) {
	counter := make([]int, 3)

	for _, n := range nums {
		counter[n]++
	}

	current, k := 0, 0

	for i := 0; i < len(nums); {
		if counter[current] == 0 {
			current++
		} else {
			nums[i] = current
			k++
			i++
		}
		if k == counter[current] {
			k = 0
			current++
		}
	}

}
