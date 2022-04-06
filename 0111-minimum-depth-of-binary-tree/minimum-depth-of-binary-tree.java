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