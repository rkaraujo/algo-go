package solution0235

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var lowestAncestor *TreeNode

	if isAncestor(root, p) && isAncestor(root, q) {
		lowestAncestor = root
	} else {
		return nil
	}

	leftAncestor := lowestCommonAncestor(root.Left, p, q)
	rightAncestor := lowestCommonAncestor(root.Right, p, q)

	if leftAncestor != nil {
		lowestAncestor = leftAncestor
	} else if rightAncestor != nil {
		lowestAncestor = rightAncestor
	}

	return lowestAncestor
}

func isAncestor(root, node *TreeNode) bool {
	if root == nil {
		return false
	}

	if root == node || root.Left == node || root.Right == node {
		return true
	}

	if node.Val < root.Val {
		return isAncestor(root.Left, node)
	} else {
		return isAncestor(root.Right, node)
	}
}
