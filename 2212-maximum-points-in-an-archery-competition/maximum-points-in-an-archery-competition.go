func maximumBobPoints(numArrows int, aliceArrows []int) []int {
    
    maxPoints := 0
    maxSlice := make([]int,12)
    
    path := make([]int,12)
    // dfs 剪枝，从分值大的开始dfs，然后判断如果后面的所有分加起来都不没办法更大的话就退出
    
    // 正向遍历
    var dfs func (curDepth,points,leftArrows int)
    dfs=func (curDepth,points,leftArrows int){
        // 终止条件
        if curDepth >11 {
            if maxPoints < points{
                if leftArrows>0{
                    path[0]+=leftArrows
                }
                maxPoints = points
                copy(maxSlice,path)
                if leftArrows>0{
                    path[0]-=leftArrows
                }
            } 
            return
        }
        // 可以选择当前节点的情况下尝试选择
        if leftArrows > aliceArrows[curDepth] {
            path[curDepth] = aliceArrows[curDepth]+1
            dfs(curDepth+1,points+curDepth,leftArrows-aliceArrows[curDepth]-1)
            path[curDepth] = 0
        }
        dfs(curDepth+1,points,leftArrows)
    }
    dfs(0,0,numArrows)
    return maxSlice
}