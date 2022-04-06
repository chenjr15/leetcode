func hammingWeight(num uint32) int {

    cnt:=0
    for num&-num !=0{
        cnt++
        num &= ^(num&-num)
    }
    return cnt
    
}