func isAlienSorted(words []string, order string) bool {
    if len(words) ==0 {
        return true
    }
    // 将顺序重映射
    orderMap := make([]int,26)
    for i:=range order{
        orderMap[order[i]-'a'] = i
    }
    // 比较字典序, 对于相同的字符跳过，知道比到一个不同的字符
    lenWords := len(words)-1
    for i:=0;i<lenWords;i++{
        
        // fmt.Println(i,words[i],words[i+1],strcmp(words[i],words[i+1],orderMap))
       if  !strcmp(words[i],words[i+1],orderMap) {
           fmt.Println(words[i],words[i+1])
           return false
       }
    }
    return true
}

func strcmp(s1,s2 string, order []int) bool{
    // len(order) == 26
    var p int
    minlen:=len(s1)
    if len(s2)<minlen{
        minlen = len(s2)
    }
    for p=0;p<minlen && s1[p] == s2[p];p++{
        
    }
    // 开始碰到不一样的字符,或者是字符串结尾
    if p==minlen{
        // 走到至少一个字符串的结尾,
        // 字符串1 比字符串2 长,返回false
        fmt.Println(s1,s2, p)

        return len(s1)<=len(s2)
    }
    // 都没有走到结尾, 且不相等，s1的字符比较小则说明是合理的,否则是不合理的
    // fmt.Println(s1,s2, s1[p],s2[p],order[s1[p]-'a'] , order[s2[p]-'a'])
    return order[s1[p]-'a'] < order[s2[p]-'a']
    
}