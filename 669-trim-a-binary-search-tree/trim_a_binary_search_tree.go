package leetcode669

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * 修剪二叉树使其元素都在指定区间内[low,high]
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 第一种需要更换节点的情况
	// 如果当前节点小于low，则用右节点取代当前节点
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	// 第二种需要更换节点的情况
	// 如果当前节点大于high，则用左节点来取代当前节点
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	// 当前节点在给定取点范围内，那么就递归修剪左右节点即可(无需更换当前节点)
	root.Right = trimBST(root.Right, low, high)
	root.Left = trimBST(root.Left, low, high)
	return root
}
