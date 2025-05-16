package tasks

func findBottomLeftValue(root *TreeNode) int {
	maxLength, count := 0, 0
	val := root.Val
	findBottom(root, maxLength, &count, &val)
	return val

}

func findBottom(node *TreeNode, maxLength int, count *int, val *int) {

	if node == nil {
		return
	}
	if node.Right == nil && node.Left == nil {
		if *count > maxLength {
			maxLength = *count
			*val = node.Val
		}
	}
	*count++
	findBottom(node.Left, maxLength, count, val)
	findBottom(node.Right, maxLength, count, val)
	*count--
}
