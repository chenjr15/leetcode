# 剑指 Offer II 025.链表中的两数相加

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/lMSNwu/](https://leetcode-cn.com/problems/lMSNwu/)

## 题目描述

<p>给定两个 <strong>非空链表</strong> <code>l1</code>和 <code>l2</code>&nbsp;来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。</p>

<p>可以假设除了数字 0 之外，这两个数字都不会以零开头。</p>

<p>&nbsp;</p>

<p><strong>示例1：</strong></p>

<p><img alt="" src="https://pic.leetcode-cn.com/1626420025-fZfzMX-image.png" style="width: 302px; " /></p>

<pre>
<strong>输入：</strong>l1 = [7,2,4,3], l2 = [5,6,4]
<strong>输出：</strong>[7,8,0,7]
</pre>

<p><strong>示例2：</strong></p>

<pre>
<strong>输入：</strong>l1 = [2,4,3], l2 = [5,6,4]
<strong>输出：</strong>[8,0,7]
</pre>

<p><strong>示例3：</strong></p>

<pre>
<strong>输入：</strong>l1 = [0], l2 = [0]
<strong>输出：</strong>[0]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>链表的长度范围为<code> [1, 100]</code></li>
	<li><code>0 &lt;= node.val &lt;= 9</code></li>
	<li>输入数据保证链表代表的数字无前导 0</li>
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>如果输入链表不能修改该如何处理？换句话说，不能对列表中的节点进行翻转。</p>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 445&nbsp;题相同：<a href="https://leetcode-cn.com/problems/add-two-numbers-ii/">https://leetcode-cn.com/problems/add-two-numbers-ii/</a></p>

## 标签

 - Stack 
 - Linked List 
 - Math 

## 代码

```golang
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
```