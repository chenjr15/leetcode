func countHillValley(nums []int) (cnt int) {
    
    // 直接暴力
    lastdiff := 0
    i:=1
    for ;i<len(nums);i++{
        if nums[i] != nums[lastdiff]{
            lastdiff = i-1
            break
        }
    }

    // lastdiff 需要和i不一样
    for i:=lastdiff +1 ;i<len(nums);{
        // 扫描下一个不一样的
        nextdiff:=i+1
        for ;nextdiff<len(nums);nextdiff++{
            if nums[i] != nums[nextdiff]{
                break
            }
        }
        // nextdiff是第一个不等于i的
        if nextdiff == len(nums){
            break
        }
        
        // 先判断是否是谷
        if nums[lastdiff]>nums[i]{
            // 可能是谷
            if nums[nextdiff] > nums[i] {
                cnt++
            }
      
            
        }else{
            // 可能是峰
            if nums[nextdiff] < nums[i] {
                // i > j 
                cnt++
            }
        }
        lastdiff = i
        i=nextdiff
    }
    return cnt
}