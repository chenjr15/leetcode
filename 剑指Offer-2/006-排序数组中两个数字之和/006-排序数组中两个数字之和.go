func twoSum(numbers []int, target int) []int {
    // 有序所以可以头尾找
    var left,right,tempSum int
    right = len(numbers)-1
    for left< right{
        tempSum = numbers[left] + numbers[right] 
        if tempSum == target{
            break
        }else if tempSum < target{
            left++
        }else{
            right--
        }
    }

    return []int{left,right}
}