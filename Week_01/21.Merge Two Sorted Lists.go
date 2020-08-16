package Week_01

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{}
	list := &head
	for {
		if l1 == nil || l2 == nil {
			break
		}
		if l1.Val < l2.Val {
			list = addNode(list, l1.Val)
			l1 = l1.Next
		} else {
			list = addNode(list, l2.Val)
			l2 = l2.Next
		}
	}

	if l1 != nil {
		list.Next = l1
	}
	if l2 != nil {
		list.Next = l2
	}

	return head.Next
}

func addNode(list *ListNode, val int) *ListNode {
	node := ListNode{
		Val: val,
	}
	list.Next = &node
	return list.Next
}
