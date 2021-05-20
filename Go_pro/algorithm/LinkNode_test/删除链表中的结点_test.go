package LinkNode_test

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Val != val {
		slow, fast = fast, fast.Next
	}
	if fast != nil {
		slow.Next = fast.Next
	}
	return head
}
