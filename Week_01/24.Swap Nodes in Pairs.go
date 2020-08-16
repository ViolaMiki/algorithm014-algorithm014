package Week_01

func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode {
		Next: head,
	}
	res := pre
	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		pre.Next = second
		first.Next = second.Next
		second.Next = first

		pre = first
		head = first.Next
	}
	return res.Next
}