/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    maxPerLevel := []int{}
    
    if root==nil{
        // 输入情况以外
        return maxPerLevel
    }
    // 遍历树 
    queue := []*TreeNode{root}
    start:=0
    // cur level 从1开始
    curLevel:=0
    nextLevelCnt:=0
    curLevelCnt:=1
    // curMax := root
    var node,curMax *TreeNode
    // 层次遍历
    for ;start<len(queue);start++{
        if curLevelCnt == 0 {
            // 换行
            maxPerLevel= append(maxPerLevel,curMax.Val)
            curMax=nil
            curLevelCnt = nextLevelCnt
            nextLevelCnt=0
            curLevel++
        }
       
        // 取出头
        node = queue[start]
        if curMax == nil || curMax.Val < node.Val{
            curMax = node
        }
        // fmt.Printf("level %d : %v\n",curLevel,node.Val)
        curLevelCnt--
        
        // 左右节点加入队列
        if node.Left!=nil{
            queue = append(queue,node.Left)
            nextLevelCnt++
        }
        if node.Right!=nil{
            queue = append(queue,node.Right)
            nextLevelCnt++
        }

    }
    if curMax!=nil{
        maxPerLevel= append(maxPerLevel,curMax.Val)
    }
    return maxPerLevel
}

