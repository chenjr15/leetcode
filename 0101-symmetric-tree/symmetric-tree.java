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
    public boolean isSymmetric(TreeNode root) {
        if (root==null){
            return true;
        }
        // 先左右翻转再判断两颗树是否一样
        return dfs(root.left,root.right);

    }

    public boolean dfs(TreeNode a,TreeNode b){
        if (a==null && b==null){
            return true;
        }
        
        if ((a==null || b == null) && a!=b){
            return false;
        }
        // 都不是null
        if (a.val!=b.val){
            return false;
        } 
        return dfs(a.left,b.right) && dfs(a.right,b.left);
    }
}