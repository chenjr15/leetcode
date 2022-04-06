# 653.两数之和 IV - 输入 BST

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/](https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/)

## 题目描述

<p>给定一个二叉搜索树 <code>root</code> 和一个目标结果 <code>k</code>，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 <code>true</code>。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/09/21/sum_tree_1.jpg" style="height: 229px; width: 400px;" />
<pre>
<strong>输入:</strong> root = [5,3,6,2,4,null,7], k = 9
<strong>输出:</strong> true
</pre>

<p><strong>示例 2：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/09/21/sum_tree_2.jpg" style="height: 229px; width: 400px;" />
<pre>
<strong>输入:</strong> root = [5,3,6,2,4,null,7], k = 28
<strong>输出:</strong> false
</pre>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li>二叉树的节点个数的范围是&nbsp;&nbsp;<code>[1, 10<sup>4</sup>]</code>.</li>
	<li><code>-10<sup>4</sup>&nbsp;&lt;= Node.val &lt;= 10<sup>4</sup></code></li>
	<li><code>root</code>&nbsp;为二叉搜索树</li>
	<li><code>-10<sup>5</sup>&nbsp;&lt;= k &lt;= 10<sup>5</sup></code></li>
</ul>

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