package leetcode98

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	last := math.MinInt64
	stack := make([]*TreeNode, 0)
	p := root
	for p != nil || len(stack) != 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if p == nil && len(stack) != 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if last >= p.Val {
				return false
			}
			last = p.Val
			p = p.Right
		}
	}
	return true

}
