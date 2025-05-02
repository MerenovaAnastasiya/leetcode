package tasks

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	minSize := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		minSize++
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.Left == nil && cur.Right == nil {
				return minSize
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return minSize
}
