func isAnagram(s string, t string) bool {
    // 字符哈希
    if len(s)!=len(t) {
        return false
    }
    hashS := [26]int{}
    hashT := [26]int{} 
    allSame:=true
    for i:=0;i<len(s);i++{
        hashS[s[i] - 'a']++
        hashT[t[i] - 'a']++
        if allSame && s[i]!=t[i] {
            allSame=false
        } 
    }
    if allSame{
        return false
    }
    
    for i:=0;i<26;i++{
        if hashS[i]!= hashT[i]{
            return false
        }
    }
    return true
}