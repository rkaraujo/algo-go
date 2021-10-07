package solution0094

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
func inorderTraversal(root *TreeNode) []int {
	result := []int{}

	if root != nil {
		left := inorderTraversal(root.Left)
		for _, val := range left {
			result = append(result, val)
		}

		result = append(result, root.Val)

		right := inorderTraversal(root.Right)
		for _, val := range right {
			result = append(result, val)
		}
	}

	return result
}
