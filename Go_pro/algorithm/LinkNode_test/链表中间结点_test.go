package LinkNode_test


func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		} else if fast.Next != nil {
			return slow.Next
		} else {
			return slow
		}
	}
}