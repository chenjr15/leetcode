func findAnagrams(s2 string, s1 string) []int {
    result:= make([]int,0)
    if len(s1)>len(s2){
        return result
    }
    // s1len := len(s1)
    // s2len := len(s2)
    // 字符串滚动哈希
    s1hash := make([]int,26)
    s2hash := make([]int,26)
    // 计算s1的哈希
    for i :=range []byte(s1){
        s1hash[s1[i] - 'a']++
        s2hash[s2[i] - 'a']++
    }
    if arrayEquals(s1hash,s2hash){
        result=append(result,0)
    }
    difflen := len(s2)-len(s1)+1
    
    for i :=1;i<difflen;i++ {
        // 加头
        s2hash[s2[i+len(s1)-1]-'a']++
        // 去尾
        s2hash[s2[i-1]-'a']--
        if arrayEquals(s1hash,s2hash){
            result=append(result,i)
        }
    }
    return result

}

func arrayEquals(a ,b []int)( equals bool){
    equals=true
    for i:= range a{
        if a[i] != b[i]{
            equals =false
            break
        }
    }
    return equals

}
