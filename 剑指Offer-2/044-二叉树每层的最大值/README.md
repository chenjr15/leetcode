# 剑指 Offer II 044.二叉树每层的最大值

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/hPov7L/](https://leetcode-cn.com/problems/hPov7L/)

## 题目描述

<p>给定一棵二叉树的根节点&nbsp;<code>root</code> ，请找出该二叉树中每一层的最大值。</p>

<p>&nbsp;</p>

<p><strong>示例1：</strong></p>

<pre>
<strong>输入: </strong>root = [1,3,2,5,3,null,9]
<strong>输出: </strong>[1,3,9]
<strong>解释:</strong>
          1
         / \
        3   2
       / \   \  
      5   3   9 
</pre>

<p><strong>示例2：</strong></p>

<pre>
<strong>输入: </strong>root = [1,2,3]
<strong>输出: </strong>[1,3]
<strong>解释:</strong>
          1
         / \
        2   3
</pre>

<p><strong>示例3：</strong></p>

<pre>
<strong>输入: </strong>root = [1]
<strong>输出: </strong>[1]
</pre>

<p><strong>示例4：</strong></p>

<pre>
<strong>输入: </strong>root = [1,null,2]
<strong>输出: </strong>[1,2]
<strong>解释:</strong>      
&nbsp;          1 
&nbsp;           \
&nbsp;            2     
</pre>

<p><strong>示例5：</strong></p>

<pre>
<strong>输入: </strong>root = []
<strong>输出: </strong>[]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>二叉树的节点个数的范围是 <code>[0,10<sup>4</sup>]</code></li>
	<li><meta charset="UTF-8" /><code>-2<sup>31</sup>&nbsp;&lt;= Node.val &lt;= 2<sup>31</sup>&nbsp;- 1</code></li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 515&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/">https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/</a></p>

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
func largestValues(root *TreeNode) []int {
    maxPerLevel := []int{}
    
    if root==nil{
        // 输入情况以外
        return maxPerLevel
    }
    // 遍历树 
    queue := []*TreeNode{root}
    start:=0
    // cur level 从1开始
    curLevel:=0
    nextLevelCnt:=0
    curLevelCnt:=1
    // curMax := root
    var node,curMax *TreeNode
    // 层次遍历
    for ;start<len(queue);start++{
        if curLevelCnt == 0 {
            // 换行
            maxPerLevel= append(maxPerLevel,curMax.Val)
            curMax=nil
            curLevelCnt = nextLevelCnt
            nextLevelCnt=0
            curLevel++
        }
       
        // 取出头
        node = queue[start]
        if curMax == nil || curMax.Val < node.Val{
            curMax = node
        }
        // fmt.Printf("level %d : %v\n",curLevel,node.Val)
        curLevelCnt--
        
        // 左右节点加入队列
        if node.Left!=nil{
            queue = append(queue,node.Left)
            nextLevelCnt++
        }
        if node.Right!=nil{
            queue = append(queue,node.Right)
            nextLevelCnt++
        }

    }
    if curMax!=nil{
        maxPerLevel= append(maxPerLevel,curMax.Val)
    }
    return maxPerLevel
}


```