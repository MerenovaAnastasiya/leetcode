package tasks

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}

	l := dfs(node.Left, res)
	r := dfs(node.Right, res)

	*res = max(*res, l+r)
	return 1 + max(l, r)
}
