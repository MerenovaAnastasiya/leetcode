package tasks

func Insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int

	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)

	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}
