func sortColors(nums []int)  {
    low:=0
    high:= len(nums)-1
    // known as `Dutch National Flag` problem
    //  http://users.monash.edu/~lloyd/tildeAlgDS/Sort/Flag/
    //  https://algorithmsandme.com/dutch-national-flag-problem/
    // 将数组分成三部分, 设数组长度为N
    // [0:low ] 0
    // [low:p ] 1
    // [p:high] unknow
    // [high:N] 2
    // 需要做的是缩减unknow, 扩张其他的区
    for p:=0;p <= high;{
        switch nums[p]{
        case 0:
            nums[low],nums[p] = nums[p],nums[low]
            low++
            p++
        case 1:
            p++
        case 2:
            nums[high],nums[p] = nums[p],nums[high]
            high--
            
        }
    }
    
}
