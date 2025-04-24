package tasks

func MajorityElement(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	n := len(nums) / 3
	m := map[int]int{}
	var res []int
	for _, v := range nums {
		if count, ok := m[v]; ok {
			count++
			m[v] = count
		} else {
			m[v] = 1
		}
		if m[v] > n {
			res = append(res, v)
		}
	}
	return res
}
