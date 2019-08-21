package leetcode234

import "github.com/chenjr15/golist/linkedlist"

type ListNode = linkedlist.ListNode

func isPalindrome(head *ListNode) bool {
	// 找中点， 反转， 比较
	// 找中点和反转可以合并起来，边找中点边反转前半部分， 找到中点之后直接一个往前一个往后比较即可
	if head == nil || head.Next == nil {
		return true
	}

	// 找中点， 快慢指针
	var fast, slow, slowNext, slowPre *ListNode
	fast = head
	slow = head
	forward := false

	for fast.Next != nil {
		fast = fast.Next

		if forward {
			// 反转链表
			slowNext = slow.Next
			slow.Next = slowPre
			slowPre = slow
			slow = slowNext
		}
		forward = !forward
	}

	fast = slow.Next
	if !forward {
		// 奇数, 中间(当前的slow位置)的跳过, slow 往前移一位
		slow = slowPre
	} else {
		slow.Next = slowPre
	}
	// 比较前后是否一样
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}

	return true
}
