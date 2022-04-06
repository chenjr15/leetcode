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