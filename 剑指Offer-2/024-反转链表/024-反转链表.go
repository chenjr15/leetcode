/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(p *ListNode) *ListNode {
    var pre,next *ListNode
    for p!=nil{
        next = p.Next
        p.Next = pre
        pre = p 
        p = next 
    }
    return pre
}