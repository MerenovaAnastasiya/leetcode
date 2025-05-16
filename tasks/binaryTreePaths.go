package tasks

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	var backTracking func(*TreeNode, []*TreeNode, string)
	backTracking = func(node *TreeNode, path []*TreeNode, str string) {
		if node == nil {
			return
		}
		path = append(path, node)

		if node.Left == nil && node.Right == nil {
			str = str + strconv.Itoa(node.Val)
			result = append(result, str)
		} else {
			str = str + strconv.Itoa(node.Val) + "->"
		}
		backTracking(node.Left, path, str)
		backTracking(node.Right, path, str)
	}

	backTracking(root, nil, "")
	return result

}
