package leetcode112

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		// 空的树是没有路径的
		return false
	}
	remain := targetSum - root.Val
	if root.Left == nil && root.Right == nil && remain == 0 {
		return true
	}
	if hasPathSum(root.Left, remain) {
		return true
	}

	return hasPathSum(root.Right, remain)
}
