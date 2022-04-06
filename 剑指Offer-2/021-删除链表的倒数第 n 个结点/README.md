# 剑指 Offer II 021.删除链表的倒数第 n 个结点

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/SLwz0R/](https://leetcode-cn.com/problems/SLwz0R/)

## 题目描述

<p>给定一个链表，删除链表的倒数第&nbsp;<code>n</code><em>&nbsp;</em>个结点，并且返回链表的头结点。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg" style="width: 542px; height: 222px;" /></p>

<pre>
<strong>输入：</strong>head = [1,2,3,4,5], n = 2
<strong>输出：</strong>[1,2,3,5]
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>head = [1], n = 1
<strong>输出：</strong>[]
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>head = [1,2], n = 1
<strong>输出：</strong>[1]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>链表中结点的数目为 <code>sz</code></li>
	<li><code>1 &lt;= sz &lt;= 30</code></li>
	<li><code>0 &lt;= Node.val &lt;= 100</code></li>
	<li><code>1 &lt;= n &lt;= sz</code></li>
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>能尝试使用一趟扫描实现吗？</p>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 19&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/">https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/</a></p>

## 标签

 - Linked List 
 - Two Pointers 

## 代码

```golang
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
```