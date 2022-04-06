# 剑指 Offer II 026.重排链表

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/LGjMqU/](https://leetcode-cn.com/problems/LGjMqU/)

## 题目描述

<p>给定一个单链表 <code>L</code><em> </em>的头节点 <code>head</code> ，单链表 <code>L</code> 表示为：</p>

<p><code>&nbsp;L<sub>0&nbsp;</sub>&rarr; L<sub>1&nbsp;</sub>&rarr; &hellip; &rarr; L<sub>n-1&nbsp;</sub>&rarr; L<sub>n&nbsp;</sub></code><br />
请将其重新排列后变为：</p>

<p><code>L<sub>0&nbsp;</sub>&rarr;&nbsp;L<sub>n&nbsp;</sub>&rarr;&nbsp;L<sub>1&nbsp;</sub>&rarr;&nbsp;L<sub>n-1&nbsp;</sub>&rarr;&nbsp;L<sub>2&nbsp;</sub>&rarr;&nbsp;L<sub>n-2&nbsp;</sub>&rarr; &hellip;</code></p>

<p>不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<p><img alt="" src="https://pic.leetcode-cn.com/1626420311-PkUiGI-image.png" style="width: 240px; " /></p>

<pre>
<strong>输入: </strong>head = [1,2,3,4]
<strong>输出: </strong>[1,4,2,3]</pre>

<p><strong>示例 2:</strong></p>

<p><img alt="" src="https://pic.leetcode-cn.com/1626420320-YUiulT-image.png" style="width: 320px; " /></p>

<pre>
<strong>输入: </strong>head = [1,2,3,4,5]
<strong>输出: </strong>[1,5,2,4,3]</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>链表的长度范围为 <code>[1, 5 * 10<sup>4</sup>]</code></li>
	<li><code>1 &lt;= node.val &lt;= 1000</code></li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 143&nbsp;题相同：<a href="https://leetcode-cn.com/problems/reorder-list/">https://leetcode-cn.com/problems/reorder-list/</a>&nbsp;</p>

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
```