package LinkNode_test


// a + (b - c) = b + (a - c)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ha, hb := headA, headB
	for ha != hb {
		if ha == nil {
			ha = headB
		} else {
			ha = ha.Next
		}
		if hb == nil {
			hb = headA
		} else {
			hb = hb.Next
		}
	}
	return ha
}

