package LinkNode_test

func reserveLink(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func newLink(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	new_1 := newLink(head)
	new_2 := reserveLink(new_1.Next)

	res := true
	first := head
	second := new_2

	for res && second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}

	new_1.Next = reserveLink(new_2)
	return true
}