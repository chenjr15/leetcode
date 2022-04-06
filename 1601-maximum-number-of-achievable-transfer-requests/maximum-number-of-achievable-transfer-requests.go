func maximumRequests(n int, requests [][]int) int {
    maxCnt:=0
    l := len(requests)
    
    satified := func (plan int) int{
        satifiedPlanCnt := 0
        buildingMap:=make([]int,20)
        // 模拟搬家计划
        for i:=0;i<l;i++{
            if plan&1 == 1{
                buildingMap[requests[i][0]]--
                buildingMap[requests[i][1]]++
                satifiedPlanCnt++
            }
            plan >>=1
        }
        // 验证是否所有的人数都是刚好的
        for _,v:= range buildingMap{
            if v !=0{
                return 0
            }
        }
        return satifiedPlanCnt
        
    }
    cur:=0
    for i :=0;i<(2<<l);i++{
        // 完全穷举
        // 判断i代表的方案是否可行，对应位数为1代表满足1
        cur = satified(i)
        if maxCnt < cur {
            maxCnt = cur
        }

    }
    return maxCnt
}