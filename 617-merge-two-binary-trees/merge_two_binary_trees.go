package leetcode617

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root2 != nil && root1 != nil {
		root1.Val += root2.Val

		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
	}
	if root1 != nil {
		return root1
	}
	if root2 != nil {
		return root2
	}
	return root1
}
