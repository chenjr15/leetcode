func validUtf8(data []int) bool {

    dataLenMask:=[]int{
        0b10000000,
        0b11100000,
        0b11110000,
        0b11111000,
    }
    correctData:= []int{
        0b00000000,
        0b11000000,
        0b11100000,
        0b11110000,
    }
    idx:= 0
    runeLen := 0
    for idx<len(data){
        // 这里应该是编码的开头
        // 首先判断这个开头是不是1来决定字节数
        runeLen=0
        for k,mask:=range dataLenMask{
            // fmt.Println(data[idx],mask&data[idx],correctData[k])
            if mask&data[idx] == correctData[k]{
                runeLen=k+1
                break
            }
        }
        // 长度为0 即未匹配上或者长度超过序列长度
        if  runeLen == 0 || ( runeLen+idx)>len(data) {
                // fmt.Println(idx,data[idx],runeLen)
            return false
        }
        // 组
        for k:=1;k<runeLen;k++{
            if data[idx+k] &0b11000000 != 0b10000000{
                // fmt.Println(idx,data[idx+k])
                return false
            }
        }
        idx+=runeLen
    }
    return true

}