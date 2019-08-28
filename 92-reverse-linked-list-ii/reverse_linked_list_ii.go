package leetcode92

import "github.com/chenjr15/golist/linkedlist"

type ListNode = linkedlist.ListNode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 讲道理这题可以直接用25的函数

	p := head
	if head == nil || head.Next == nil {
		return head
	}
	if m == 1 {
		return reverseN(p, n-m+1)
	}
	pos := 1
	// 循环结束后p指向要反转的节点的前一个
	for pos < m-1 && p != nil {
		pos++
		p = p.Next
	}

	pre := p
	pre.Next = reverseN(p.Next, n-m+1)
	return head

}

// 部分逆序，reverseLength 指定要逆序的长度，然返回新的头节点
func reverseN(head *ListNode, reverseLength int) (newhead *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	var next, pre *ListNode
	// 逐个反转
	for reverseLength > 0 && p != nil {
		next = p.Next
		p.Next = pre
		pre = p
		p = next
		reverseLength--
	}

	head.Next = p
	return pre

}
