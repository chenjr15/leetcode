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
    int[] nums;
    public TreeNode sortedArrayToBST(int[] nums) {
        // 找中点作为根，然后递归处理左右节点。
        this.nums = nums;
        return dfs(0,nums.length-1);
    }

    public TreeNode dfs(int left,int right){
        TreeNode node =null;
        if(left>right){
            return null;
        }
        int mid = (left+right)/2;
        node = new TreeNode(nums[mid]);
        node.left = dfs(left,mid-1);
        node.right = dfs(mid+1,right);
        return node;
    }   
}