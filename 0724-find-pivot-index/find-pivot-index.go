func pivotIndex(nums []int) int {
    // 前缀和
    // 不能用二分，直接从左往右遍历
    psum:=0
    sum := 0
    for i := range nums{
        sum+=nums[i]
    }
    for i:=range nums{
        // 利用左右相等的特性，直接将判断条件改为，i左边的前缀和+i右边的前缀和+i的值 == 总和
        // 左右前缀和相等因此可以改成 i左边前缀和*2 +nums[i] == 总和
        if psum+psum+nums[i] == sum{
            return i
        }
        psum+=nums[i]
    }
    return -1
}