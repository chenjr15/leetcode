/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// rotateRight Linking the linked list into a ring, then you can rotate it.
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    length := 1
    
    // 找尾
    p := head
    for ;p.Next!=nil;p = p.Next{
        length++
    }
    // 接环
    p.Next = head
    // 旋转
    k %= length
    k = length-k
    for p= head;k>0;k--{
        head = head.Next
    }
    // 断环
    for p = head;p.Next != head ;p = p.Next {
    }
    p.Next = nil
    
    return head
    
}
