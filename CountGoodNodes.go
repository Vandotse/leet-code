/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}
	res := 0
	if node.Val >= maxVal {
		res = 1
		maxVal = node.Val

	}
	if node.Val > maxVal {
		maxVal = node.Val
	}
	res += dfs(node.Left, maxVal)
	res += dfs(node.Right, maxVal)
	return res
}