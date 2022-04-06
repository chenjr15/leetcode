func threeSum(nums []int) [][]int {
    lenNums := len(nums)
    results:=make([][]int,0,100)
    if lenNums<3{
        return results
    }
    // 固定一个元素，搜索剩下两个，变成一个twosum问题
    // 将元素排序，加快后面的twoSum搜索
    sort.Ints(nums)
    var pb,pc,b,c,tempSum int
    
    for pa,a := range nums{
        if a>0{
            // 那么后面一定也大于0了，所有大于0的数字加起来不会等于0 的
            break
        }
        if pa>0 && nums[pa] == nums[pa-1]{
            continue
        }
        // a+b+c = 0 ==> b+c = -a
        pb,pc=pa+1,lenNums-1
        // 双指针搜索,注意这里会导致每一个都会搜索一次
        // 这里要注意因为有重复元素，所以要注意去重
        a=-a
        for pb<pc {
            b = nums[pb]
            c = nums[pc]
            tempSum = b+c
            
            if tempSum == a{
                results = append(results,[]int{-a,b,c})
                pc--
                pb++
                for pb<pc&& nums[pb] == b {
                    pb++
                }  
                
                for pb<pc && nums[pc] == c {
                    pc-- 
                }
            }else if tempSum< a {
                // for pb<pc&& nums[pb] == b {
                    pb++
                // }  
            }else{
                // for pb<pc && nums[pc] == c {
                    pc-- 
                // }
            }
        }
    }
    return results
}