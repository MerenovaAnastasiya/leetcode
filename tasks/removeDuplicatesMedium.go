package tasks

func RemoveDuplicatesMedium(nums []int) int {
	k := 1
	current := nums[0]
	countSame := 1
	counter := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != current {
			countSame = 1
			counter++
			nums[k] = nums[i]
			k++
			current = nums[i]
		} else {
			countSame++
			if countSame == 2 {
				counter++
				nums[k] = nums[i]
				k++
				current = nums[i]
			}
		}
	}
	return counter

}
