package leetcode94

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归版本
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var inorder func(t *TreeNode)
	inorder = func(t *TreeNode) {

		if t.Left != nil {
			inorder(t.Left)
		}
		result = append(result, t.Val)
		if t.Right != nil {
			inorder(t.Right)
		}

	}
	if root == nil {
		return []int{}
	}
	result = make([]int, 0, 10)
	inorder(root)
	return result

}
