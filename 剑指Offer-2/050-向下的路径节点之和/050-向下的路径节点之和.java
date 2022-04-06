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