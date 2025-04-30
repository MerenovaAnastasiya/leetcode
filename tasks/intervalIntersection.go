package tasks

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var intersections [][]int
	i, j := 0, 0

	for i < len(firstList) && j < len(secondList) {
		if firstList[i][1] >= secondList[j][0] && firstList[i][0] <= secondList[j][1] {
			intersections = append(intersections, []int{
				max(firstList[i][0], secondList[j][0]),
				min(firstList[i][1], secondList[j][1])})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return intersections
}
