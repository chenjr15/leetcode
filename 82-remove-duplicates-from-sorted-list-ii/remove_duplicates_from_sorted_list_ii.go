package leetcode82

import "github.com/chenjr15/golist/linkedlist"

type ListNode = linkedlist.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	// 首先这是排序了的
	// 判断重复只需要比较当前元素和下一个元素是否一样
	// 可以采取碰到重复的就把所有重复的删掉
	// 即直接快进到下一个非当前重复元素，然后重复

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
		for p.Next != nil && p.Next.Val == p.Val {
			p = p.Next
		}
		if pre == nil {
			// 上一个为空说明至今未找到非重复节点的作为头节点
			head = p.Next
		} else {
			pre.Next = p.Next
		}
		p = p.Next

	}
	return head

}
