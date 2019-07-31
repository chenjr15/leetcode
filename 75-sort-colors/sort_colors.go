func sortColors(nums []int)  {
    // 大概算减治法
    low:=0
    // 分步计算, 遍历超过一次
    // 把所有0移到数组开头
    for p:=0;p < len(nums);p++{
        if nums[p] ==0 {
                nums[low],nums[p] = nums[p],nums[low]
            low ++
        }
    }
    // 把所有1移到数组开头
    for p:=low;p < len(nums);p++{
        if nums[p] ==1 {
                nums[low],nums[p] = nums[p],nums[low]
            low ++
        }
    }
    
}
