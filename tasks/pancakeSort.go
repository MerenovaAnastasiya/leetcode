package tasks

func PancakeSort(arr []int) []int {
	left, right := 0, len(arr)-1
	var result []int
	for left < right {
		maxIndex := left

		for i := left; i <= right; i++ {
			if arr[i] > arr[maxIndex] {
				// находим индекс максимального
				maxIndex = i
			}
		}
		if maxIndex != right {
			if maxIndex != left {
				flip(arr, maxIndex+1)
				result = append(result, maxIndex+1)
			}
			flip(arr, right+1)
			result = append(result, right+1)
		}
		right--
	}

	return result

}

func flip(arr []int, k int) {
	for i := 0; i < k/2; i++ {
		arr[i], arr[k-1-i] = arr[k-1-i], arr[i]
	}
}
