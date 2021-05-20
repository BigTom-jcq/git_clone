package LinkNode_test

func swapPairs(head *LinkNode) *LinkNode {
	k := &LinkNode{}
	k.next = head
	p := k
	for head != nil && head.next != nil {
		next := head.next
		head.next = next.next
		next.next = head
		p.next = next

		p = head
		head = head.next
	}
	return k.next
}