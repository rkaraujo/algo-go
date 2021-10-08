package solution0102

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
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	if root != nil {
		idx_level := 0
		cur_level := []*TreeNode{}
		cur_level = append(cur_level, root)

		for len(cur_level) > 0 {
			next_level := []*TreeNode{}
			result[idx_level] = make([]int, 0)

			for _, node := range cur_level {
				result[idx_level] = append(result[idx_level], node.Val)

				if node.Left != nil {
					next_level = append(next_level, node.Left)
				}
				if node.Right != nil {
					next_level = append(next_level, node.Right)
				}
			}

			cur_level = next_level
			idx_level++
		}
	}

	return result
}
