func singleNumber(nums []int) int {
    result :=0
    // 按二进制逐位计算
    // 对于每一位来说，如果把所有数字对应的1加起来的话得到和为sum。
    // 因为除了目标元素以为，其他元素都出现三次，如果那个元素对应的二进制为1 那么一定使sum加3 如果是0则无影响。
    // 同样的，目标元素会使sum+1。其他元素对sum的影响可以用过对3区域来去除，只剩下目标元素带来影响的1 或者0
    sum:=0
    for i:=0;i<32;i++{
        sum =0 
        for _,v:=range nums{
            sum+=((v>>i)&1)
        }
        // 去除该元素的影响
        result |= ((sum%3)<<i  )
    }
    return int(int32(result))
    
}