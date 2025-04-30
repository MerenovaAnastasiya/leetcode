package tasks

func longestMountain(arr []int) int {
	// 2,1,4,7,3,2,5

	if len(arr) < 3 {
		return 0
	}
	maxLength := 0

	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] && arr[i] > arr[i-1] {
			left := i
			right := i
			for left > 0 && arr[left] > arr[left-1] {
				left--
			}
			for right < len(arr) && arr[right] > arr[right+1] {
				right++
			}
			maxLength = max(maxLength, right-left+1)
		}
	}
	return maxLength

}
