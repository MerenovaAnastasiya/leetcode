package tasks

func SearchRange(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	result := make([]int, 0)
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid + 1
		}
		if nums[mid] > target {
			end = mid - 1
		}

		if nums[mid] == target {
			if nums[start] == target && nums[end] == target {
				return append(result, start, end)
			} else {
				if nums[start] != target {
					start++
				}
				if nums[end] != target {
					end--
				}
			}
		}
	}
	if len(result) == 0 {
		return []int{-1, -1}
	} else {
		return result
	}
}
