package solution0203

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	// find head
	for head != nil && head.Val == val {
		head = head.Next
	}

	// remove elements
	cur := head
	for cur != nil {
		next := cur.Next
		for next != nil && next.Val == val {
			next = next.Next
		}
		cur.Next = next

		cur = cur.Next
	}

	return head
}
