func subarraySum(nums []int, k int) int {
    // 前缀和
    prefixCount:=make(map[int]int)
    result :=0
    psum:=0
    // 计算[0,i]的情况
    prefixCount[0]=1
    for i:=range nums{
        // [0,i] 的前缀和
        psum+=nums[i]
        result += prefixCount[ (psum - k )]
        prefixCount[psum]++
    }
    return result

}