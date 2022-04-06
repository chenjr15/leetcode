/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public TreeNode inorderSuccessor(TreeNode root, TreeNode p) {
        // 利用二叉搜索树的性质，如果
        TreeNode result = null;
        // 否则那结果就是他的父节点
        while(root!=null ){
            // 题目保证了值唯一，因此不用考虑值相等。
            if (root.val>p.val){
                result = root;
                root=root.left;
            }else{
                root=root.right;
            }
        }

        return result;
    }
}