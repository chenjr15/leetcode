func countBits(n int) []int {

    // 按二进制拆数字，后面的用前面的来计算
    result := make([]int,n+1)
    //当前位数
  
    for i:=1 ;i<=n ;i++{
        result[i] += result[i>>1] + i &1
        // result[i] = oneCount
    }
    return result
}