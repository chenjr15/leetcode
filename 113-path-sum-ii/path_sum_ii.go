package leetcode113

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		// 空的树是没有路径的
		return ret
	}
	path := make([]int, 0)

	var dfs func(p *TreeNode, r int)
	dfs = func(p *TreeNode, r int) {
		if p == nil {
			return
		}
		path = append(path, p.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		r -= p.Val
		if p.Left == nil && p.Right == nil {
			// fmt.Println("path:",path)
			// 叶节点
			if r == 0 {
				// 复制一份path
				path_copy := make([]int, len(path))
				copy(path_copy, path)
				ret = append(ret, path_copy)
			}
			return
		}
		dfs(p.Left, r)
		dfs(p.Right, r)
	}
	dfs(root, targetSum)
	return ret
}
