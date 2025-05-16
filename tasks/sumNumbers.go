package tasks

func SumNumbers(root *TreeNode) int {
	return findSum(root, 0)
}

func findSum(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	sum = sum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return sum
	}
	return findSum(node.Left, sum) + findSum(node.Right, sum)
}
