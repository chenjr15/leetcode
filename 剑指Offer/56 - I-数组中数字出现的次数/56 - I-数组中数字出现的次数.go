func singleNumbers(nums []int) []int {
    var sum,x1,x2 int
    // 求出sum=x1^x2
    for _,n:=range nums{
        sum^= n
    }
    // 求最低有效位
    lsb:=sum & -sum
    for _,n:=range nums{
        // 按lsb分类异或
        if n & lsb !=0{
            x1^=n
        }else {
            x2^=n
        }
    }
    return []int{x1,x2}

}