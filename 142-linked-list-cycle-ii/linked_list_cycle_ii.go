package leetcode142
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * 代码在这
 * https://github.com/chenjr15/golist/blob/master/linkedlist/linkedlist.go
 * 链表找环：快慢指针，相遇既是有环，相遇后一个在相遇点，一个在表头，以相等的速度走，再次相遇即是环的入口
 */
func detectCycle(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return nil
	}
	forward := true
	slow := l
	var p *ListNode
	for fast := l.Next; fast != nil; fast = fast.Next {
		if fast.Next == slow.Next {
			p = slow.Next
			break
		}
		if forward {
			slow = slow.Next
		}
		forward = !forward

	}
	for slow = l; p != nil; {
		if slow == p {
			return slow
		}
		p = p.Next
		slow = slow.Next

	}
	return nil
}
