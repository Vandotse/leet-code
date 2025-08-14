/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var resp int

func diameterOfBinaryTree(root *TreeNode) int {
	resp = 0

	dfs(root)
	return resp
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	resp = max(resp, right+left)
	return 1 + max(left, right)
}
