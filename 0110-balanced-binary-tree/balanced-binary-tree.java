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
    boolean balanced = true;
    public boolean isBalanced(TreeNode root) {
        // 肯定是递归
        // 然后返回最深
        getHeight(root);
        return balanced;
    }

    public int getHeight(TreeNode node){
        if (node == null){
            return 0;
        }
        int leftHeight = getHeight(node.left);
        if (leftHeight==-1){
            return -1;
        }
        int rightHeight = getHeight(node.right);
        if (rightHeight==-1){
            return -1;
        }
        if(Math.abs(leftHeight-rightHeight)>1){
            // 退出
            balanced =  false;
            return -1;
        }
       
        return Math.max(leftHeight,rightHeight)+1;
    }
}