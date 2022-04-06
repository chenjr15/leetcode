# 剑指 Offer II 050.向下的路径节点之和

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/6eUYwP/](https://leetcode-cn.com/problems/6eUYwP/)

## 题目描述

<p>给定一个二叉树的根节点 <code>root</code>&nbsp;，和一个整数 <code>targetSum</code> ，求该二叉树里节点值之和等于 <code>targetSum</code> 的 <strong>路径</strong> 的数目。</p>

<p><strong>路径</strong> 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<p><img src="https://assets.leetcode.com/uploads/2021/04/09/pathsum3-1-tree.jpg" style="width: 452px; " /></p>

<pre>
<strong>输入：</strong>root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
<strong>输出：</strong>3
<strong>解释：</strong>和等于 8 的路径有 3 条，如图所示。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
<strong>输出：</strong>3
</pre>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li>二叉树的节点个数的范围是 <code>[0,1000]</code></li>
	<li><meta charset="UTF-8" /><code>-10<sup><span style="font-size: 9.449999809265137px;">9</span></sup>&nbsp;&lt;= Node.val &lt;= 10<sup><span style="font-size: 9.449999809265137px;">9</span></sup></code>&nbsp;</li>
	<li><code>-1000&nbsp;&lt;= targetSum&nbsp;&lt;= 1000</code>&nbsp;</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 437&nbsp;题相同：<a href="https://leetcode-cn.com/problems/path-sum-iii/">https://leetcode-cn.com/problems/path-sum-iii/</a></p>

## 标签

 - Tree 
 - Depth-First Search 
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
    ArrayList<Integer> prefixSum;
    int target;
    int cnt;
    public int pathSum(TreeNode root, int targetSum) {
        prefixSum = new ArrayList<Integer>();
        prefixSum.add(0);
        target = targetSum;
        cnt=0;
        // dfs前缀和
        dfs(root);
        return cnt;

    }
    
    public void dfs(TreeNode node){
        if (node == null){
            return;
        }
    
        int curSum = prefixSum.get(prefixSum.size()-1);
        curSum+=node.val;
        // 找一遍prefixsum,以当前节点为终点的区间和
        for(int i = 0;i<prefixSum.size();++i){
            if (curSum-prefixSum.get(i) == target){
                ++cnt;
            }
        }
        prefixSum.add(curSum);
        
        // System.out.println("Node:"+node.val+prefixSum.toString());
        dfs(node.left);
        // prefixSum.set(prefixSum.size()-1,curSum);
        dfs(node.right);
        prefixSum.remove(prefixSum.size()-1);
        return ;

    }
}
```