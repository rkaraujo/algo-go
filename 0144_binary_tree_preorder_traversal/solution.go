package solution0144

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
func preorderTraversal(root *TreeNode) []int {
	result := []int{}

	if root != nil {
		result = append(result, root.Val)

		left := preorderTraversal(root.Left)
		for _, val := range left {
			result = append(result, val)
		}

		right := preorderTraversal(root.Right)
		for _, val := range right {
			result = append(result, val)
		}
	}

	return result
}
