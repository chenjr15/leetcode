func twoSum(numbers []int, target int) []int {
    // 双指针 
    // 先让指针指向两边端点，如果加起来的和比较大那就 右边的往左缩，否则往右缩
    left:=0
    right:=len(numbers)-1
    sum:=0
    for left<right{
        sum =  numbers[left]+numbers[right]
        if sum == target{
            return []int{left+1,right+1}
        }else if sum>target{
            right--
        }else{
            left++
        }
    }
    return []int{-1,-1}
}