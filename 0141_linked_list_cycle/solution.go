package solution0141

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
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	visited := make(map[*ListNode]bool)

	cur := head
	for cur != nil {
		if visited[cur] {
			return true
		}
		visited[cur] = true
		cur = cur.Next
	}

	return false
}
