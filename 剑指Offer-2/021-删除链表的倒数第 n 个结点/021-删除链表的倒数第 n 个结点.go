/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil{
        return nil
    }
    fast := head
    slow := head
    
    // 先走n步
     i:=0
    for;i<n&& fast.Next!=nil;i++{
       
        fast = fast.Next
        
    }
    // 走不满n步的情况下要把头节点删掉
    if i!=n {
        return head.Next
    }

    // 开始遍历,找到节点末尾, 最后fast应该指向倒数第二个节点,slow指向倒数n+1个节点
    for fast.Next!=nil{
       
        fast = fast.Next
        slow = slow.Next
        
    }
    if slow.Next!=nil{
          // 将slow指向的节点
        // slow = slow.Next 
        toRemove := slow.Next
        slow.Next = toRemove.Next
    }
  

    return head
}