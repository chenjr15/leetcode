/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    // 把第一条链的末尾和连到第一条链的末尾
    // 如果他们有重合的话,就从第二条链开始找环的入口,那就肯定是交点
    // 然后把尾指针断开就好了
    tail:=headA
    for tail!=nil&&tail.Next!=nil{
        tail=tail.Next
    }
    // 把末尾和headA接起来
    tail.Next= headA

    fast:= headB
    slow:= headB
    cross :=headB
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
        fast = headB
        for fast!=slow {
            fast=fast.Next
            slow = slow.Next
        }
        cross= slow
    }
    tail.Next = nil
    return cross

}