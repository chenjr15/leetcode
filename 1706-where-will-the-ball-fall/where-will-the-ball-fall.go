func findBall(grid [][]int) []int {
    if len(grid)==0 || len(grid[0])==0{
        return []int{}
    }
    n:=len(grid[0])
    cur:=grid[0][0]
    answer:=make([]int,n)
    for i:=0;i<n;i++{
        cur=i;
        // i 是第i个球
        for _,line :=range grid{
            if line[cur] == 1{
                // 导向右边，判断右边的格子
                // 在最右那格或者右边是导向左边的
                if cur==n-1 || line[cur+1] == -1{
                    // 卡住
                    cur=-1
                    break
                    
                }
                cur++
            }else{
                // 导向左边，判断左边的格子
                if cur==0 || line[cur-1] == 1{
                    cur=-1
                    
                    break
                }
                cur--
            }
        }
        answer[i] = cur
    }
    return answer
}