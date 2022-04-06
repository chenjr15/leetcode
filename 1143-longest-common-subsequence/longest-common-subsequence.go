func longestCommonSubsequence(text1 string, text2 string) int {
    if len(text1) ==0 || len(text2) == 0 {
        return 0
    }
    len1 := len(text1)+1
    len2 := len(text2)+1
    dp := make([][]int,len1)
    for i:=range dp{
        dp[i] = make([]int,len2)
    }

    // dp[i][j]  是text1 以i-1下标结尾和text2 以j-1下标结尾的字符串最长的公共子序列
    // 故 dp[len(text1)][len(text2)] 即为答案

    // 那么base case 就是
    // dp[0][*] = 0 
    // dp[*][0] = 0
    // dp[1][1] = (text1[0] == text2[0]) (这个条件可以不要，直接用转移方程推出来)
    
    // 转移方程：
    // if text1[i-1] == text2[j-1]  
    //  dp[i][j] = dp[i-1][j-1]+1
    // else 
    // dp[i][j] = max(dp[i-1][j],dp[i][j-1])

    // text1end
    for i:=1;i<len1;i++{
        
        for j:=1;j<len2;j++{
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1]+1
            }else{
                // dp[i][j] = max(dp[i-1][j],dp[i][j-1])
                if dp[i][j-1]>dp[i-1][j]{
                    dp[i][j] = dp[i][j-1]

                }else{
                    dp[i][j] = dp[i-1][j]
                }
            }
        }
    }

    return dp[len1-1][len2-1]
}

// func max(a,b int ) int{
//     if a>b{
//         return a
//     }
//     return b

// }