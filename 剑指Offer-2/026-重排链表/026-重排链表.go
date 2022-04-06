/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    // 可以看到，该方法就是将整个元素拆分为前后两个部分，然后讲后面的部分反转再插到前面的部分中。
    // 因此需要做的是用快慢指针找到中点，然后断开链表，再将其合并
    // 如果空间能用 O(N)的话，也可以用一个数组存起来，不过这里要交换节点，所以应该不能用这个
    slow:=head
    fast:=head 
    
    for fast!=nil &&fast.Next!=nil{
        fast = fast.Next.Next
        slow=slow.Next
    }
    if fast == slow {
        // 长度小于3，无需交换
        return 
    }
    // 注意中点是往右偏的，在奇数的情况下slow刚好指向的是中点，中点后的元素是我们要的
    // 如果是偶数的话，中点有两个，slow指向的是右边的中点，我们只需要插入右边中点以后的节点即可
    // 因此该遍历方法会自动指向中间偏右的位置，无需特殊判断
    pright := reverse(slow.Next)
    pleft := head
    // 断开链表
    slow.Next = nil
    next:= pright
    // right 肯定小于等于left，right肯定先走完或者同时走完
    for pleft!=nil && pright!=nil {
        // 把右边的间隔插入左边的
        next = pleft.Next
        pleft.Next = pright
        pright=pright.Next
        pleft.Next.Next = next
        pleft=next
    }
    return 
}
func reverse(p*ListNode) (r *ListNode){
    var pre,next *ListNode
    for p!=nil{
        next = p.Next
        p.Next = pre
        pre = p 
        p = next 
    }
    return pre
}