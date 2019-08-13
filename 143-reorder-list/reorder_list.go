package leetcode143

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return

	}
	slow, fast := head, head
	forward := false
	for fast != nil {
		fast = fast.Next
		if forward {
			slow = slow.Next

		}
		forward = !forward

	}
	p := slow.Next
	slow.Next = nil

	// Reverse
	var last *ListNode
	var next *ListNode
	if p == nil {
		return

	}

	for p.Next != nil {
		next = p.Next
		p.Next = last
		last = p
		p = next

	}
	p.Next = last

	l := head
	// Merge
	for l != nil && p != nil {
		next = l.Next
		l.Next = p
		l = l.Next
		p = p.Next
		l.Next = next
		l = next

	}

}
