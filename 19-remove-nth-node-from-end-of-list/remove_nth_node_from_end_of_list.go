/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if n<1{
        return head
    }
    pre := head
    p:= head
    // forward the fast pointer n step first
    for n>0   {
        p=p.Next
        n--
    }
    // fast pointer reachs the end which  means that n equals the length of list, remove head
    if p == nil {
        return head.Next
    }
    // move until the end 
    for p.Next!=nil{
        p=p.Next
        pre=pre.Next
    }
    // remove nth node
    pre.Next = pre.Next.Next
    return head
    
}
