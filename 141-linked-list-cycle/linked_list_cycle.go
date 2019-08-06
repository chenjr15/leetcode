/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * 代码在这
 * https://github.com/chenjr15/golist/blob/master/linkedlist/linkedlist.go
 */
 func hasCycle(l *ListNode) bool {
	if l == nil || l.Next == nil {
	 return false
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
		 return slow != nil
	 }
	 p = p.Next
	 slow = slow.Next

 }
 return false
}