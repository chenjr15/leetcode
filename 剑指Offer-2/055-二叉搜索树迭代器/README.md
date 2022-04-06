# 剑指 Offer II 055.二叉搜索树迭代器

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/kTOapQ/](https://leetcode-cn.com/problems/kTOapQ/)

## 题目描述

<p>实现一个二叉搜索树迭代器类<code>BSTIterator</code> ，表示一个按中序遍历二叉搜索树（BST）的迭代器：</p>

<div class="original__bRMd">
<div>
<ul>
	<li><code>BSTIterator(TreeNode root)</code> 初始化 <code>BSTIterator</code> 类的一个对象。BST 的根节点 <code>root</code> 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。</li>
	<li><code>boolean hasNext()</code> 如果向指针右侧遍历存在数字，则返回 <code>true</code> ；否则返回 <code>false</code> 。</li>
	<li><code>int next()</code>将指针向右移动，然后返回指针处的数字。</li>
</ul>

<p>注意，指针初始化为一个不存在于 BST 中的数字，所以对 <code>next()</code> 的首次调用将返回 BST 中的最小元素。</p>
</div>
</div>

<p>可以假设&nbsp;<code>next()</code>&nbsp;调用总是有效的，也就是说，当调用 <code>next()</code>&nbsp;时，BST 的中序遍历中至少存在一个下一个数字。</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2018/12/25/bst-tree.png" style="width: 189px; height: 178px;" /></p>

<pre>
<strong>输入</strong>
inputs = [&quot;BSTIterator&quot;, &quot;next&quot;, &quot;next&quot;, &quot;hasNext&quot;, &quot;next&quot;, &quot;hasNext&quot;, &quot;next&quot;, &quot;hasNext&quot;, &quot;next&quot;, &quot;hasNext&quot;]
inputs = [[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
<strong>输出</strong>
[null, 3, 7, true, 9, true, 15, true, 20, false]

<strong>解释</strong>
BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next();    // 返回 3
bSTIterator.next();    // 返回 7
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 9
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 15
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 20
bSTIterator.hasNext(); // 返回 False
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>树中节点的数目在范围 <code>[1, 10<sup>5</sup>]</code> 内</li>
	<li><code>0 &lt;= Node.val &lt;= 10<sup>6</sup></code></li>
	<li>最多调用 <code>10<sup>5</sup></code> 次 <code>hasNext</code> 和 <code>next</code> 操作</li>
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong></p>

<ul>
	<li>你可以设计一个满足下述条件的解决方案吗？<code>next()</code> 和 <code>hasNext()</code> 操作均摊时间复杂度为 <code>O(1)</code> ，并使用 <code>O(h)</code> 内存。其中 <code>h</code> 是树的高度。</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 173&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/binary-search-tree-iterator/">https://leetcode-cn.com/problems/binary-search-tree-iterator/</a></p>

## 标签

 - Stack 
 - Tree 
 - Design 
 - Binary Search Tree 
 - Binary Tree 
 - Iterator 

## 代码

```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    stack []*TreeNode
    
}
func stack2str (nodes []*TreeNode)string{
    if len(nodes) ==0 {
        return ""
    }
    builder := strings.Builder{}
    for i:=0;i<len(nodes);i++{
        builder.WriteString(strconv.Itoa(i))
        builder.WriteRune(':')
        builder.WriteString(strconv.Itoa(nodes[i].Val))
        builder.WriteRune(',')
    }
    // builder.WriteString(strconv.Itoa(nodes[len(nodes)-1].Val))
    return builder.String()
}

func Constructor(root *TreeNode) BSTIterator {
    stack:= []*TreeNode{}
    p := root
    // --- 标准非递归中序遍历 ---
    // for p !=nil || len(stack)!=0{
    //     // 一直往左遍历到当前的最左节点，并全部入栈
    //     for p!=nil{
    //         stack = append(stack,p)
    //         p = p.Left
    //     }
    //     if len(stack)>0 {
    //         // 出栈
    //         p = stack[len(stack)-1]
    //         stack = stack[:len(stack)-1]
    //         // 出栈的时候就是要访问的时候
    //         // visit(p)
    //         fmt.Println(p)
    //         p = p.Right
    //     }
    // }
    // 一直往左遍历到当前的最左节点，并全部入栈
    for p!=nil{
            stack = append(stack,p)
            p = p.Left
    }
    return BSTIterator{stack}

    
}


func (this *BSTIterator) Next() (nextval int) {
    // 题目说保证有效，暂时不管stack是否为空
    // fmt.Println("before next",stack2str(this.stack))
    p := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    nextval = p.Val
    p = p.Right
    for p!=nil{
        this.stack = append(this.stack,p)
        p = p.Left
    }
    // fmt.Println("after ",nextval,stack2str(this.stack))
    return nextval
}   


func (this *BSTIterator) HasNext() bool {
    return len(this.stack)!=0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
```