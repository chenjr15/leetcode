func countSubstrings(s string) int {
    // 遍历中心点
    // 单个长度的可以直接用长度
    total:=len(s)
    var left,right int
    // 单个的中心点
    for i:=range s {
        left = i -1
        right = i+1 
        // fmt.Println(i,left,right, left>=0 && right <len(s)&& s[left] == s[right])
        for left>=0 && right <len(s)&& s[left] == s[right]{
            total++
            left--
            right++
        }
    }
    // 双个的中心点
    for i:=0;i<len(s)-1;i++ {
        
        left = i
        right = i+1 
        // fmt.Println(i,left,right, left>=0 && right <len(s)&& s[left] == s[right])
        for left>=0 && right <len(s)&& s[left] == s[right]{
            total++
            left--
            right++
        }
    }
    return total 
}