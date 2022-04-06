# 剑指 Offer II 056.二叉搜索树中两个节点之和

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/opLdQZ/](https://leetcode-cn.com/problems/opLdQZ/)

## 题目描述

<p>给定一个二叉搜索树的 <strong>根节点</strong> <code>root</code>&nbsp;和一个整数 <code>k</code> , 请判断该二叉搜索树中是否存在两个节点它们的值之和等于 <code>k</code> 。假设二叉搜索树中节点的值均唯一。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入: </strong>root =<strong> </strong>[8,6,10,5,7,9,11], k = 12
<strong>输出: </strong>true
<strong>解释: </strong>节点 5 和节点 7 之和等于 12
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入: </strong>root =<strong> </strong>[8,6,10,5,7,9,11], k = 22
<strong>输出: </strong>false
<strong>解释: </strong>不存在两个节点值之和为 22 的节点
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>二叉树的节点个数的范围是&nbsp;&nbsp;<code>[1, 10<sup>4</sup>]</code>.</li>
	<li><code>-10<sup>4</sup>&nbsp;&lt;= Node.val &lt;= 10<sup>4</sup></code></li>
	<li><code>root</code>&nbsp;为二叉搜索树</li>
	<li><code>-10<sup>5</sup>&nbsp;&lt;= k &lt;= 10<sup>5</sup></code></li>
</ul>

<p>&nbsp;</p>

<p>注意：本题与主站 653 题相同：&nbsp;<a href="https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/">https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/</a></p>

## 标签

 - Tree 
 - Depth-First Search 
 - Breadth-First Search 
 - Binary Search Tree 
 - Hash Table 
 - Two Pointers 
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
func findTarget(root *TreeNode, k int) bool {
    // 尝试模仿双指针
    var sleft,sright []*TreeNode
    // fmt.Println(sright)
    pleft := root
    pright := root
    for pleft!=nil{
        sleft= append(sleft,pleft)
        pleft = pleft.Left
    }
    goRight := func ()(r *TreeNode){
        if len(sleft)==0{
            return nil
        }
        pleft = sleft[len(sleft)-1]
        sleft = sleft[:len(sleft)-1]
        r = pleft
        pleft = pleft.Right
        for pleft!=nil{
            sleft= append(sleft,pleft)
            pleft = pleft.Left
        }
        return r
    }
    for pright!=nil{
        sright= append(sright,pright)
        pright = pright.Right
    }
    // 从右往左走
    goLeft := func ()(r *TreeNode){
        if len(sright)==0{
            return nil
        }
        pright = sright[len(sright)-1]
        sright = sright[:len(sright)-1]
        r = pright
        pright = pright.Left
        for pright!=nil{
            sright= append(sright,pright)
            pright = pright.Right
        }
        return r
    }

    // 左中右，中序遍历的非递归实现
    // for {
    //     p:= goRight()
    //     if p==nil{
    //         break
    //     }
    //     fmt.Println(p.Val)
    // }
    // for {
    //     p:= goLeft()
    //     if p==nil{
    //         break
    //     }
    //     fmt.Println(p.Val)
    // }
    var l,r *TreeNode
    
    l=goRight()
    r=goLeft()
    for l!=r{
        sum:=l.Val+r.Val
        if sum == k{
            return true
        }else if sum>k{
            r = goLeft()
        }else{
            l = goRight()
        }
    }
  

    return false

}


```