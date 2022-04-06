func minSubArrayLen(target int, nums []int) int {
    minlen:= 0
    nlen:=len(nums)
    // 尝试滑动窗口
    // 首先对范围有一个明确的定义，子数组的范围就是[left,right],那么sum=nums[left,...,right],
    // curlen = right-left +1, left<=right < nlen 
    var left,right,sum,curlen int
    sum = nums[0]
    curlen=1
    for left <= right {
        // 往右扩直到范围内的数组和大于target
        for right<nlen-1 && sum < target{
            right++
            sum += nums[right]
        }
        
        for left <= right && sum>= target{
            curlen= right-left+1 
            if minlen == 0 || curlen<minlen{
                minlen = curlen
            }
            // 往左缩
            sum -= nums[left]
            left++
        }
        // 跳出的关键，右边扩张到最后一个，且左边已经收缩完了，就可以退出了
        if right==nlen-1{
            break
        }
        
    }
    return minlen

}