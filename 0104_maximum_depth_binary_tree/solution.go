package solution0104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDepthLeft := 1 + maxDepth(root.Left)
	maxDepthRight := 1 + maxDepth(root.Right)

	if maxDepthLeft > maxDepthRight {
		return maxDepthLeft
	}
	return maxDepthRight
}
