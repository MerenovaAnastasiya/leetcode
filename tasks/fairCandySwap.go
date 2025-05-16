package tasks

import "slices"

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	alice := sumArr(aliceSizes)
	bob := sumArr(bobSizes)

	// sumA - a + b = sumB - b + a
	// sumA - sumB - 2(a - b) = 0
	// sumA - sumB = 2 * (a - b)
	// (sumA - sumB)/2 = a - b
	// diff = a - b

	diff := (alice - bob) / 2

	slices.Sort(aliceSizes)

	for _, b := range bobSizes {
		if a := binarySearch(aliceSizes, b+diff); a != -1 {
			return []int{a, b}
		}
	}
	return []int{}

}

func sumArr(arr []int) int {
	arrSum := 0
	for _, n := range arr {
		arrSum += n
	}
	return arrSum
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		med := l + (r-l)/2
		if arr[med] == target {
			return target
		} else if arr[med] > target {
			r = med - 1
		} else {
			l = med + 1
		}
	}
	return -1
}
