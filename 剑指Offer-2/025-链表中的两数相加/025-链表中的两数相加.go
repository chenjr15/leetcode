/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // 反转链表再相加
    l1 = reverse(l1)
    l2 = reverse(l2)
    resultHead:=&ListNode{}
    l3 :=resultHead
    // 尝试相加
    carry:=0
    for l1!=nil && l2!=nil{
        temp:= &ListNode{Val:l1.Val+l2.Val+carry}
        carry = temp.Val/10
        temp.Val %=10
        l3.Next = temp
        l3 = temp
        l1=l1.Next
        l2=l2.Next
    }
    
    // 接下去就是高位
    if l2 !=nil{
        l1=l2
    }
    for l1!=nil{
        temp:= &ListNode{Val:l1.Val+carry}
        carry = temp.Val/10
        temp.Val %=10
        l3.Next = temp
        l3 = temp
        l1=l1.Next
    }
    if carry !=0{
        l3.Next = &ListNode{Val:carry}
    }
    return reverse(resultHead.Next)
}

func reverse(p *ListNode) *ListNode{
    var pre,next *ListNode
    for p!=nil&&p.Next!=nil{
        next = p.Next
        p.Next = pre
        pre = p 
        p = next
    }
    if p!=nil{
        p.Next = pre
    }
    return p
}