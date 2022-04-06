/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    fast:= head
    slow:=head
    cross :=head
    cross = nil
    for fast!=nil &&fast.Next!=nil{
        fast = fast.Next.Next 
        slow = slow.Next
        if fast == slow{
            cross = slow
            break
        }
    }
    if cross != nil{
        // 有环, 开始找环
        fast = head
        for fast!=slow {
            fast=fast.Next
            slow = slow.Next
        }
        return slow
    }else{
        return nil
    }
}