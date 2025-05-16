package tasks

func isBalancedTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	flag := true
	findLength(root, &flag)
	return flag
}

func findLength(node *TreeNode, flag *bool) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}

	l := findLength(node.Left, flag)
	r := findLength(node.Right, flag)
	if abs(l-r) > 1 {
		*flag = false
	}

	return 1 + max(l, r)
}
