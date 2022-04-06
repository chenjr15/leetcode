# 剑指 Offer II 027.回文链表

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/aMhZSa/](https://leetcode-cn.com/problems/aMhZSa/)

## 题目描述

<p>给定一个链表的 <strong>头节点&nbsp;</strong><code>head</code><strong>&nbsp;，</strong>请判断其是否为回文链表。</p>

<p>如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<p><strong><img alt="" src="https://pic.leetcode-cn.com/1626421737-LjXceN-image.png" /></strong></p>

<pre>
<strong>输入:</strong> head = [1,2,3,3,2,1]
<strong>输出:</strong> true</pre>

<p><strong>示例 2：</strong></p>

<p><strong><img alt="" src="https://pic.leetcode-cn.com/1626422231-wgvnWh-image.png" style="width: 138px; height: 62px;" /></strong></p>

<pre>
<strong>输入:</strong> head = [1,2]
<strong>输出:</strong> false
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>链表 L 的长度范围为 <code>[1, 10<sup><span style="font-size: 9.449999809265137px;">5</span></sup>]</code></li>
	<li><code>0&nbsp;&lt;= node.val &lt;= 9</code></li>
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>能否用&nbsp;O(n) 时间复杂度和 O(1) 空间复杂度解决此题？</p>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 234&nbsp;题相同：<a href="https://leetcode-cn.com/problems/palindrome-linked-list/">https://leetcode-cn.com/problems/palindrome-linked-list/</a></p>

## 标签

 - Stack 
 - Recursion 
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
```