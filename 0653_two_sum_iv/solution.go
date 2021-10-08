package solution0653

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
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	return findTargetFromCurNode(root, root, k)
}

func findTargetFromCurNode(root *TreeNode, curNode *TreeNode, k int) bool {
	if curNode == nil {
		return false
	}

	remaining := k - curNode.Val

	otherNode := searchBST(root, curNode, remaining)
	if otherNode != nil {
		return true
	}

	hasTargetLeft := findTargetFromCurNode(root, curNode.Left, k)
	if hasTargetLeft {
		return true
	}

	hasTargetRight := findTargetFromCurNode(root, curNode.Right, k)
	return hasTargetRight
}

func searchBST(root *TreeNode, curNode *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val && root != curNode {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, curNode, val)
	} else {
		return searchBST(root.Right, curNode, val)
	}
}
