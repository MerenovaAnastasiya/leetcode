package tasks

import (
	"sort"
)

func merge(intervals [][]int) [][]int {

	if len(intervals) == 1 {
		return intervals
	}

	// отсортировали по левой точке каждого отрезка
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// [[1,3],[2,6],[8,10],[15,18]]
	result := [][]int{intervals[0]}

	for _, interval := range intervals {
		lastMerged := result[len(result)-1]
		if interval[0] <= lastMerged[1] {
			lastMerged[1] = max(lastMerged[1], interval[1])
		} else {
			result = append(result, interval)
		}
	}
	return result

}
