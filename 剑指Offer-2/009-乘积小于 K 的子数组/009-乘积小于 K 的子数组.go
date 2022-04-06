func numSubarrayProductLessThanK(nums []int, target int) int {
    // 应该是滑动窗口,拿上一题过来套

    arrCnt:= 0
    nlen:=len(nums)
    // 尝试滑动窗口
    // 首先对范围有一个明确的定义，子数组的范围就是[left,right),那么sum=nums[left,...,right-1],
    // curlen = right-left, left<right <= nlen 
    var left,right,product int
    right=1
    product = nums[0]

    for left < right {
        // 往右扩直到范围内的乘积大于等于target
        for right<nlen && product < target{
            product *= nums[right]
            right++
        }
        // right 是第一个不满足条件的位置，因此
        // fmt.Println(left,right,product)
        // 得到以left开头的，right-1结尾的子数组乘积满足条件
        arrCnt+= right-left-1
        if product < target{
            arrCnt++
        }
        if left == nlen{
            break
        }
        product/=nums[left]
        left++
        // for left < right && sum>= target{
        //     curlen= right-left+1 
        //     if minlen == 0 || curlen<minlen{
        //         minlen = curlen
        //     }
        //     // 往左缩
        //     sum -= nums[left]
        //     left++
        // }
        // // 跳出的关键，右边扩张到最后一个，且左边已经收缩完了，就可以退出了
        // if right==nlen-1{
        //     break
        // }
        
    }
    return arrCnt

}