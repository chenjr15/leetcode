# 剑指 Offer II 045.二叉树最底层最左边的值

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/LwUNpT/](https://leetcode-cn.com/problems/LwUNpT/)

## 题目描述

<p>给定一个二叉树的 <strong>根节点</strong> <code>root</code>，请找出该二叉树的&nbsp;<strong>最底层&nbsp;最左边&nbsp;</strong>节点的值。</p>

<p>假设二叉树中至少有一个节点。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<p><img src="https://assets.leetcode.com/uploads/2020/12/14/tree1.jpg" style="width: 182px; " /></p>

<pre>
<strong>输入: </strong>root = [2,1,3]
<strong>输出: </strong>1
</pre>

<p><strong>示例 2: </strong></p>

<p><img src="https://assets.leetcode.com/uploads/2020/12/14/tree2.jpg" style="width: 242px; " /><strong> </strong></p>

<pre>
<strong>输入: </strong>[1,2,3,4,null,5,6,null,null,7]
<strong>输出: </strong>7
</pre>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li>二叉树的节点个数的范围是 <code>[1,10<sup>4</sup>]</code></li>
	<li><meta charset="UTF-8" /><code>-2<sup>31</sup>&nbsp;&lt;= Node.val &lt;= 2<sup>31</sup>&nbsp;- 1</code>&nbsp;</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 513&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/find-bottom-left-tree-value/">https://leetcode-cn.com/problems/find-bottom-left-tree-value/</a></p>

## 标签

 - Tree 
 - Depth-First Search 
 - Breadth-First Search 
 - Binary Tree 

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
func findBottomLeftValue(root *TreeNode) int {
    // 很显然可以用bfs 层序遍历，但是这次考虑试一下用dfs
    // 用一个maxDepth 记录遍历到的最深的深度
    maxDepth:= 0 
    // 再用一个leftmostVal 记录最深处的值，每次左边先遍历，用的后序遍历
    leftmostVal :=0
    var dfs func (node *TreeNode,curDepth int) 
    dfs = func (node *TreeNode,curDepth int)  {
        if node == nil{
            return 
        }
        if curDepth > maxDepth {
            maxDepth = curDepth
            leftmostVal = node.Val
        }
        dfs(node.Left,curDepth+1)
        dfs(node.Right,curDepth+1)
    }

    dfs(root,1)

    return leftmostVal
}
```