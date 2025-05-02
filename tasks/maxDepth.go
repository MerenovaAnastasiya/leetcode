package tasks

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	maxLength := 0

	for len(queue) > 0 {
		maxLength++
		levelLength := len(queue)

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

	return maxLength
}
