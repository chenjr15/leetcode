package leetcode24

import "github.com/chenjr15/golist/linkedlist"

type ListNode = linkedlist.ListNode

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseN(head, 2)
	pre := head
	p := head.Next
	for p != nil && p.Next != nil {
		// 保证前面的指向新的头
		pre.Next = reverseN(p, 2)
		pre = p
		p = p.Next
	}
	return newHead

}

// 部分逆序，reverseLength 指定要逆序的长度，然会新的头节点
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
