func pivotIndex(nums []int) int {
    // 前缀和+二分查找左边界
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
    
    // for i:=range nums{
    //     if psum[i] == (psum[len(nums)] - psum[i+1]){
    //         return i
    //     }
    // }
    
    // left,right,mid := 0,len(nums),len(nums)>>1
    // lsum,rsum := 0,0
    // // result:=-1
    // for left<right{
    //     mid=left+(right-left)>>1
    //     lsum = psum[mid]
    //     rsum = psum[len(nums)] - psum[mid+1]
    //     fmt.Println(left,right,mid,lsum,rsum)
    //     if lsum == rsum{
    //         right=mid
    //         // result=mid
    //         // break
    //     }else if lsum<right{
    //         left=mid+1
    //     }else{
    //         right=mid
    //     }
    // }
    // if lsum != rsum{
    //     return -1
    // }
    // return left

}