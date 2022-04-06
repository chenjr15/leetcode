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
    int sum=0;
    public int sumNumbers(TreeNode root) {
        if (root == null){
            return 0;
        }
        dfs(root,0);
        return sum;
    }

    public void dfs(TreeNode node,int curSum){
        if (node == null){
            return ;
        }
        
        if (node.left ==null && node.right == null){
            // 到达根节点
            this.sum += curSum*10 + node.val;
            return;
        }
        dfs(node.left,curSum*10+node.val);
        dfs(node.right,curSum*10+node.val);
    }


}