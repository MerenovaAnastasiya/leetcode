package tasks

func removeElement(nums []int, val int) int {
	k := 0
	j := 0

	for k < len(nums) {
		var element = nums[k]
		if element != val {
			nums[j] = nums[k]
			j++
		}
		k++
	}
	return j
}
