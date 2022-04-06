func singleNumber(nums []int) int {
    result := 0
    sum:=0
    for i:=0;i<32;i++{
        sum=0
        for _,v:=range nums{
            sum+= (v>>i)&1
        }
        sum%=3
        result |= (sum&1)<<i

    }

    return int(int32(result))
 
}