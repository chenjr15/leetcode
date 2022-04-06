func maximumDifference(nums []int) int {
    if len(nums) ==0{
        return -1
    }
    min :=nums[0]

    diff :=0
    maxDiff :=-1
    for _,n := range nums{
        diff = n-min
        if n<min{
            min = n
        }
        if diff >0 && diff > maxDiff{
            maxDiff = diff
        }
    }
    return maxDiff

}