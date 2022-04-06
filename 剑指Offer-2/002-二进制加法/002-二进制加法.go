func addBinary(a string, b string) string {
    lenA := len(a)
    lenB := len(b)
    // 二进制加法
    maxLen := lenA
    if lenB>maxLen{
        maxLen = lenB
    }
    result:= make([]byte,maxLen,maxLen+1)
    var carry byte
    i:=0
    for ;i<maxLen;i++{
        if i <lenA{
            carry+= a[lenA - i -1] - '0'
        }
        if i <lenB{
            carry+= b[lenB - i -1] - '0'
        }
        if carry & 1 != 0{
            result[i] = '1'
        }else{
            result[i] = '0'
        }
        carry>>=1
    }
    if carry>0{
        result = append(result,'1')
    }
    // reverse
    l:=len(result)
    for i,j:=0,l-1;i<j;{
        result[i],result[j] = result[j],result[i]
        i++
        j--
    }
    return string(result)

}