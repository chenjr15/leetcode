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
    public TreeNode pruneTree(TreeNode root) {

        // dfs 剪枝
        if (isAllZero(root)){
            return null;
        }
        return root;

    }

    public boolean isAllZero(TreeNode node){
        if (node==null){
            return true;
        }
        boolean leftZero = isAllZero(node.left);
        boolean rightZero = isAllZero(node.right);
        if (leftZero){
            node.left = null;
        }
        if (rightZero){
            node.right = null;
        }
        
        return node.val == 0 && leftZero && rightZero;
    }
}