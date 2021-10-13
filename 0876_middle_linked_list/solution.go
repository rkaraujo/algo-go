package solution0876

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
func middleNode(head *ListNode) *ListNode {
	middle := head
	end := head

	for end != nil && end.Next != nil {
		middle = middle.Next
		// end runs the double
		end = end.Next.Next
	}

	return middle
}
