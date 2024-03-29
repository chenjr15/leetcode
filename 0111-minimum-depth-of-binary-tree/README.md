# 111.二叉树的最小深度

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)

## 题目描述

<p>给定一个二叉树，找出其最小深度。</p>

<p>最小深度是从根节点到最近叶子节点的最短路径上的节点数量。</p>

<p><strong>说明：</strong>叶子节点是指没有子节点的节点。</p>

<p> </p>

<p><strong>示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/12/ex_depth.jpg" style="width: 432px; height: 302px;" />
<pre>
<strong>输入：</strong>root = [3,9,20,null,null,15,7]
<strong>输出：</strong>2
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>root = [2,null,3,null,4,null,5,null,6]
<strong>输出：</strong>5
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li>树中节点数的范围在 <code>[0, 10<sup>5</sup>]</code> 内</li>
	<li><code>-1000 <= Node.val <= 1000</code></li>
</ul>

## 标签

 - Tree 
 - Depth-First Search 
 - Breadth-First Search 
 - Binary Tree 

## 代码

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    int minDepth = Integer.MAX_VALUE;
    public int minDepth(TreeNode root) {
        if(root==null){
            return 0;
        }
        minDepth(root,1);
        return minDepth;
    }
    public void minDepth(TreeNode node,int curDepth){
        if(node==null){
            return ;
        }
        if(curDepth>=minDepth){
            return;
        }
        if (node.left==null&&node.right == null){
            // 叶节点
            if(curDepth<minDepth){
                minDepth = curDepth;
                return;
            }
        }

        minDepth(node.left,curDepth+1);
        minDepth(node.right,curDepth+1);

        return  ;
    }
}
```