package leetcode83

import "github.com/chenjr15/golist/linkedlist"

type ListNode = linkedlist.ListNode

func deleteDuplicates(head *ListNode) *ListNode {

	// 82题简化版

	// 先把特殊情况去除
	if head == nil || head.Next == nil {
		return head
	}
	p := head

	var pre *ListNode = nil
	for p != nil && p.Next != nil {
		// 非重复， 直接下一个
		if p.Val != p.Next.Val {
			pre = p
			p = p.Next
			continue
		}
		// 当前为重复节点
		pre = p
		for p.Next != nil && p.Next.Val == p.Val {
			p = p.Next
		}
		pre.Next = p.Next

		p = p.Next

	}
	return head

}
