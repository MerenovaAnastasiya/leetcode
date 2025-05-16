package tasks

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	inorder(root, &res)
	return res
}

func inorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, res)
	*res = append(*res, node.Val)
	inorder(node.Right, res)
}
