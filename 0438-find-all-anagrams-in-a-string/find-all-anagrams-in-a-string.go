func findAnagrams(s string, p string) []int {
    result:= make([]int,0)
    if len(p)>len(s){
        return result
    }

    // 字符串滚动哈希
    patternHash := make([]int,26)
    strHash := make([]int,26)
    // 计算pattern的哈希
    for i :=range []byte(p){
        patternHash[p[i] - 'a']++
        strHash[s[i] - 'a']++
    }
    if arrayEquals(patternHash,strHash){
        result=append(result,0)
    }
    difflen := len(s)-len(p)+1
    
    for i :=1;i<difflen;i++ {
        // 加头
        strHash[s[i+len(p)-1]-'a']++
        // 去尾
        strHash[s[i-1]-'a']--
        if arrayEquals(patternHash,strHash){
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