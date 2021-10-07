package solution0145

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
func postorderTraversal(root *TreeNode) []int {
	result := []int{}

	if root != nil {
		left := postorderTraversal(root.Left)
		for _, val := range left {
			result = append(result, val)
		}

		right := postorderTraversal(root.Right)
		for _, val := range right {
			result = append(result, val)
		}

		result = append(result, root.Val)
	}

	return result
}
