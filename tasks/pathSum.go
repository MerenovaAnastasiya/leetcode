package tasks

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var path []int
	findPath(root, targetSum, &res, path)
	return res

}

func findPath(node *TreeNode, targetSum int, res *[][]int, path []int) {
	if node == nil {
		return
	}
	path = append(path, node.Val)
	if node.Left == nil && node.Right == nil && node.Val == targetSum {
		*res = append(*res, append([]int{}, path...))
	}
	findPath(node.Left, targetSum-node.Val, res, path)
	findPath(node.Right, targetSum-node.Val, res, path)
	path = path[:len(path)-1]
}
