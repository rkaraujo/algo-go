package solution0101

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
func isSymmetric(root *TreeNode) bool {
	// init left and right
	left := []*TreeNode{}
	left = appendNode(left, root.Left)

	right := []*TreeNode{}
	right = appendNode(right, root.Right)

	// check if left level is symmetric to right level
	for len(left) > 0 && len(right) > 0 {
		if len(right) != len(left) {
			return false
		}
		for i := 0; i < len(left); i++ {
			if left[i].Val != right[len(right)-1-i].Val {
				return false
			}
		}

		// init next level
		next_left := []*TreeNode{}
		isAllNil := true
		for _, node := range left {
			if node.Left != nil || node.Right != nil {
				isAllNil = false
			}
			next_left = appendNode(next_left, node.Left)
			next_left = appendNode(next_left, node.Right)
		}
		if isAllNil {
			left = []*TreeNode{}
		} else {
			left = next_left
		}

		next_right := []*TreeNode{}
		isAllNil = true
		for _, node := range right {
			if node.Left != nil || node.Right != nil {
				isAllNil = false
			}
			next_right = appendNode(next_right, node.Left)
			next_right = appendNode(next_right, node.Right)
		}
		if isAllNil {
			right = []*TreeNode{}
		} else {
			right = next_right
		}
	}

	return true
}

func appendNode(nodes []*TreeNode, node *TreeNode) []*TreeNode {
	if node != nil {
		nodes = append(nodes, node)
	} else {
		nodes = append(nodes, &TreeNode{-999, nil, nil})
	}
	return nodes
}
