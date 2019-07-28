package mergelists

// ListNode typedef
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	if l1 == nil && l2 == nil {
		return l1
	}
	// 判断有一个输入为空的情况
	if l1 == nil && l2 != nil {
		return l2
	}
	if l2 == nil && l1 != nil {
		return l1
	}
	// 确定头结点
	if l1.Val < l2.Val {
		l3 = l1
		l1 = l1.Next

	} else {
		l3 = l2
		l2 = l2.Next
	}
	r := l3
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l3.Next = l1
			l1 = l1.Next
			l3 = l3.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
			l3 = l3.Next
		}
	}
	if l1 != nil {
		l3.Next = l1
	}
	if l2 != nil {
		l3.Next = l2
	}

	return r

}
