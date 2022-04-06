/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    // 1. 考虑存到数组里面 T: O(N) S: O(N) 
    // 2. 找中点然后反转再比较  T: O(N) S: O(1) 
    // 找中点要找中间偏右的，即如果是奇数则找中间，如果是偶数，则右边的中点
    if head == nil {
        return true 
    }
    mid := center(head)
    mid = reverse(mid)
    // 接下去按照较短的那个部分来找
    
    for mid!=nil{
        if mid.Val!= head.Val{
            return false
        }
        mid=mid.Next
        head=head.Next
    }
    return true
}

func center(head *ListNode) (*ListNode) {
    // 该遍历方法会自动指向中间偏右(偶数长度情况下)的位置，奇数情况下是正中间

    slow:=head
    fast:=head 
    
    for fast!=nil && fast.Next!=nil{
        fast = fast.Next.Next
        slow=slow.Next
    }
    return slow
    
   
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