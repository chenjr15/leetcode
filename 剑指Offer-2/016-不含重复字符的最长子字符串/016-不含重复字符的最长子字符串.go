func lengthOfLongestSubstring(s string) int {
    // 最近的一个位置
    hash :=[128]int{}
    start := 0
    result := 0
    pos:=0
    for i:=range []byte(s){
        if pos = hash[s[i]];pos>start{
            // 当前窗口中存在该字符
            if (i-start)>result{
                result = i-start
            }
            // 更新字符串起点，为当前字符上一次出现的位置
            start = pos
           
        }
         // 更新或加入哈希表
        hash[s[i]]=i+1
    }
    // 答案在结尾
    if (len(s)-start)>result{
        result = len(s)-start
    }

    return result
   
}