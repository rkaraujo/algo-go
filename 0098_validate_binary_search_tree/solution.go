package solution0098

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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isParentBiggerThanAll(root, root.Left) || !isParentSmallerThanAll(root, root.Right) {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}

func isParentBiggerThanAll(parent, subtree *TreeNode) bool {
	if subtree == nil {
		return true
	}
	if parent.Val > subtree.Val && isParentBiggerThanAll(parent, subtree.Left) && isParentBiggerThanAll(parent, subtree.Right) {
		return true
	}
	return false
}

func isParentSmallerThanAll(parent, subtree *TreeNode) bool {
	if subtree == nil {
		return true
	}
	if parent.Val < subtree.Val && isParentSmallerThanAll(parent, subtree.Right) && isParentSmallerThanAll(parent, subtree.Left) {
		return true
	}
	return false
}
