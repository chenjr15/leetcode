func addDigits(num int) int {
    if num==0{
        return 0
    }
    r:=num%9
    if r==0{
        r=9
    }
    return r 

}