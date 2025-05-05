package tasks

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	var result []int

	for len(queue) > 0 {
		levelLength := len(queue)
		result = append(result, queue[levelLength-1].Val)
		for i := 0; i < levelLength; i++ {
			current := queue[0]
			queue = queue[1:]
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}
	return result
}
