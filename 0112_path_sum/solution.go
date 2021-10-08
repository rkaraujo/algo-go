package solution0112

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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	remaining := targetSum - root.Val

	// if no more remaining and is leaf
	if remaining == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil {
		hasPathSumLeft := hasPathSum(root.Left, remaining)
		if hasPathSumLeft {
			return true
		}
	}

	if root.Right != nil {
		hasPathSumRight := hasPathSum(root.Right, remaining)
		return hasPathSumRight
	}

	return false
}
