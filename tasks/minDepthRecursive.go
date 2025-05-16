package tasks

func minDepthRec(root *TreeNode) int {
	return minDepth(root)
}

func recursiveMinDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := recursiveMinDepth(node.Left)
	right := recursiveMinDepth(node.Right)

	return min(left, right) + 1
}
