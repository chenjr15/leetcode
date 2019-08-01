package leetcode86

import (
	"github.com/chenjr15/golist/linkedlist"
)

func partition(head *linkedlist.ListNode, x int) *linkedlist.ListNode {
	var left *linkedlist.ListNode = nil
	var right *linkedlist.ListNode = nil
	left = &linkedlist.ListNode{0, nil}
	right = &linkedlist.ListNode{0, nil}
	lHead := left
	rHead := right

	for p := head; p != nil; p = p.Next {
		switch {

		case p.Val < x:
			left.Next = p
			left = left.Next
		case p.Val >= x:
			right.Next = p
			right = right.Next

		}

	}
	if right != nil {
		right.Next = nil
	}
	if lHead.Next == nil {
		return rHead.Next

	} else {
		left.Next = rHead.Next
		return lHead.Next
	}
}
