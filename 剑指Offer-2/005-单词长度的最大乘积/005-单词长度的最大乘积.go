func maxProduct(words []string) int {
    // 判断是否有相同的字母的方法：考虑到字符的范围只有小写字母，用位运算来相与即可
    // 但是还是不知道如何匹配出最长的乘积。
    // 朴素算法就是穷举
    lenwords := len(words)
    bitmaps:= make([]int,lenwords)
    for i,word:=range words{
        for j:=range word{
            bitmaps[i] |= 1<< (word[j] - 'a') 
        }
    }
    // 最大的乘积
    result := 0
    tempProduct:=0
    for i,bitmap := range bitmaps{
        for j:=i+1;j<lenwords;j++{
            // 两个单词无相交字符
            if bitmap & bitmaps[j] == 0 {
                tempProduct = len(words[i]) * len(words[j])
                if tempProduct > result {
                    result = tempProduct
                }
            }
        }
    }
    return result

}