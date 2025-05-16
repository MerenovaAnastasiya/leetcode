package tasks

import "slices"

func smallestFromLeaf(root *TreeNode) string {

	var backTracking func(*TreeNode, []*TreeNode)
	curr := ""
	minStr := ""
	backTracking = func(node *TreeNode, path []*TreeNode) {
		if node == nil {
			return
		}

		curr += string(rune('a' + node.Val))

		if node.Left == nil && node.Right == nil {
			res := rvrs(curr)
			if minStr == "" || res < minStr {
				minStr = res
			}
		}
		path = append(path, node)

		backTracking(node.Left, path)
		backTracking(node.Right, path)
		path = path[:len(path)-1]
		curr = curr[:len(curr)-1]
	}
	backTracking(root, []*TreeNode{})
	return minStr
}

func rvrs(s string) string {
	r := []rune(s)
	slices.Reverse(r)
	return string(r)
}
