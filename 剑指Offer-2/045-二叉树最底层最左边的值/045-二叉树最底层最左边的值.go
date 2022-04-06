/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    // 很显然可以用bfs 层序遍历，但是这次考虑试一下用dfs
    // 用一个maxDepth 记录遍历到的最深的深度
    maxDepth:= 0 
    // 再用一个leftmostVal 记录最深处的值，每次左边先遍历，用的后序遍历
    leftmostVal :=0
    var dfs func (node *TreeNode,curDepth int) 
    dfs = func (node *TreeNode,curDepth int)  {
        if node == nil{
            return 
        }
        if curDepth > maxDepth {
            maxDepth = curDepth
            leftmostVal = node.Val
        }
        dfs(node.Left,curDepth+1)
        dfs(node.Right,curDepth+1)
    }

    dfs(root,1)

    return leftmostVal
}