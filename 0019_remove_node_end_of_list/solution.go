package solution0019

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := ListNode{0, head}
	cur := &dummy
	end := &dummy

	// create a n distance between cur and end
	for i := 0; i < n; i++ {
		end = end.Next
	}

	// go to end of list
	for end.Next != nil {
		cur = cur.Next
		end = end.Next
	}

	// remove next
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}

	return dummy.Next
}
