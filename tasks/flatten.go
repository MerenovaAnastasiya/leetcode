package tasks

func flatten(root *TreeNode) {
	linkedList(root)
}

func linkedList(node *TreeNode) {
	if node == nil {
		return
	}
	linkedList(node.Left)
	linkedList(node.Right)

	left := node.Left

	if left != nil {
		cur := left
		for cur.Right != nil {
			cur = cur.Right
		}
		cur.Right = node.Right
		node.Right = left
		node.Left = nil
	}
}
