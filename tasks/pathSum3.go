package tasks

func PathSum3(root *TreeNode, targetSum int) int {
	count := 0
	var path []int
	rec(root, targetSum, &count, &path)
	return count
}

func rec(node *TreeNode, targetSum int, count *int, path *[]int) {
	if node == nil {
		return
	}
	*path = append(*path, node.Val)

	rec(node.Left, targetSum, count, path)
	rec(node.Right, targetSum, count, path)

	sum := 0
	for i := len(*path); i >= 0; i-- {
		sum += (*path)[i]
		if sum == targetSum {
			*count++
		}
	}
	*path = (*path)[:len(*path)-1]
}
