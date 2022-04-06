func maximumTop(nums []int, k int) int {
    // k 为零直接返回第一个数据
    if k==0{
        if len(nums)>0{
            return nums[0]
        }else{
            return -1
        }
    }
    // k=1 栈顶只能是第k+1 个
    if k==1 {
        if len(nums)>k{
            return nums[k]
        }else{
            return -1
        }
    }
    if len(nums)==1{
        if k%2 == 1 {
            return -1
        } else{
            return nums[0]
        }
    }
    
    // k>2 那就可以是前k-1个和第k+1个里面最大的，减1是为了能够放一个回去
    prefix := k-1
    if prefix >len(nums){
        prefix = len(nums)
    }
    max:=0
    for i:=1;i<prefix;i++{
        if nums[i]>nums[max]{
            max = i
        }
    }
    // 如果能够全部出栈，即第不放回去，那么判断最后一个是否比较大
    if k<len(nums) && nums[k] > nums[max]{
        max=k
    }
    return nums[max]

}