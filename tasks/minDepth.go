package tasks

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	minLength := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		minLength++
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			if current.Left == nil && current.Right == nil {
				return minLength
			}

			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}
	return minLength
}
