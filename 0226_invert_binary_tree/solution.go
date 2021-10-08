package solution0226

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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// invert subtrees
	right := invertTree(root.Right)
	left := invertTree(root.Left)

	// invert current node
	root.Left = right
	root.Right = left

	return root
}
