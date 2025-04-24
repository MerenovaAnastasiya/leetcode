package tasks

import "fmt"

func SummaryRanges(nums []int) []string {
	// Input nums = [0,2,3,4,6,8,9]
	// Output ["0","2->4","6","8->9"]
	start := 0
	if len(nums) == 1 {
		return []string{fmt.Sprint(nums[start])}
	}
	var result []string

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			if i == start+1 {
				result = append(result, fmt.Sprint(nums[start]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
			}
			start = i
		}
		if i == len(nums)-1 {
			if start == len(nums)-1 {
				result = append(result, fmt.Sprint(nums[start]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[i]))
			}
		}
	}
	return result
}
