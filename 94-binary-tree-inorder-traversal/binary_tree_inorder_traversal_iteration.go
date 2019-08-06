package leetcode94

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 迭代版本
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	p := root
	for p != nil || len(stack) != 0 {

		if p != nil {
			// 当前结点入栈
			stack = append(stack, p)
			// 往左边走
			p = p.Left
		}

		// 如果左边为空,且栈非空
		if p == nil && len(stack) != 0 {
			// 出栈
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 访问当前结点
			result = append(result, p.Val)
			// 往右走
			p = p.Right
		}
		// 如果当前结点为空且栈为空的话就是遍历完成了，应该直接返回

	}
	return result

}
