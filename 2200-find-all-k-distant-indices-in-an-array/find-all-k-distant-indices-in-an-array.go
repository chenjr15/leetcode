func findKDistantIndices(nums []int, key int, k int) []int {
    result:=[]int{}
    if len(nums)<1{
        return result
    }
    idxArr:= make([]bool, len(nums))
    for i,n := range nums{
        if n==key{
            // 将前后的k个位置都加进去
            p := i-k
            if p <0 {
                p = 0
            }
            end:=i+k+1
            if end>len(nums){
                end=len(nums)
            }
            for p<end{
                idxArr[p] = true
                p++
            }
        }
    }
    for i,ok:=range idxArr{
        if ok {
            result=append(result,i)
        }
    }
    return result
}