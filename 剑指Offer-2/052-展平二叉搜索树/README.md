# 剑指 Offer II 052.展平二叉搜索树

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/NYBBNL/](https://leetcode-cn.com/problems/NYBBNL/)

## 题目描述

<p>给你一棵二叉搜索树，请&nbsp;<strong>按中序遍历</strong> 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2020/11/17/ex1.jpg" style="width: 600px; height: 350px;" /></p>

<pre>
<strong>输入：</strong>root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
<strong>输出：</strong>[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
</pre>

<p><strong>示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2020/11/17/ex2.jpg" style="width: 300px; height: 114px;" /></p>

<pre>
<strong>输入：</strong>root = [5,1,7]
<strong>输出：</strong>[1,null,5,null,7]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>树中节点数的取值范围是 <code>[1, 100]</code></li>
	<li><code>0 &lt;= Node.val &lt;= 1000</code></li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 897&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/increasing-order-search-tree/">https://leetcode-cn.com/problems/increasing-order-search-tree/</a></p>

## 标签

 - Stack 
 - Tree 
 - Depth-First Search 
 - Binary Search Tree 
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
    ArrayList<TreeNode> nodeList;
    public TreeNode increasingBST(TreeNode root) {
        if (root==null){
            return root;
        }
        nodeList = new ArrayList<TreeNode>();
        inOrder(root);
        root = nodeList.get(0);
        root.left = null;
        TreeNode p = root;
        for(int i = 1;i<nodeList.size();++i){
            p = nodeList.get(i);
            p.left = null;
            root.right = p;
            root= p;

        }

        return nodeList.get(0);

    }
    public void inOrder(TreeNode node){
        if (node==null){
            return;
        }
        inOrder(node.left);
        nodeList.add(node);
        inOrder(node.right);
    }
}
```