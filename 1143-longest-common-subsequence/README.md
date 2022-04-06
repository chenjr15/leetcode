# 1143.最长公共子序列

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/longest-common-subsequence/](https://leetcode-cn.com/problems/longest-common-subsequence/)

## 题目描述

<p>给定两个字符串 <code>text1</code> 和 <code>text2</code>，返回这两个字符串的最长 <strong>公共子序列</strong> 的长度。如果不存在 <strong>公共子序列</strong> ，返回 <code>0</code> 。</p>

<p>一个字符串的 <strong>子序列</strong><em> </em>是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。</p>

<ul>
	<li>例如，<code>"ace"</code> 是 <code>"abcde"</code> 的子序列，但 <code>"aec"</code> 不是 <code>"abcde"</code> 的子序列。</li>
</ul>

<p>两个字符串的 <strong>公共子序列</strong> 是这两个字符串所共同拥有的子序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>text1 = "abcde", text2 = "ace" 
<strong>输出：</strong>3  
<strong>解释：</strong>最长公共子序列是 "ace" ，它的长度为 3 。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>text1 = "abc", text2 = "abc"
<strong>输出：</strong>3
<strong>解释：</strong>最长公共子序列是 "abc" ，它的长度为 3 。
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>text1 = "abc", text2 = "def"
<strong>输出：</strong>0
<strong>解释：</strong>两个字符串没有公共子序列，返回 0 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= text1.length, text2.length <= 1000</code></li>
	<li><code>text1</code> 和 <code>text2</code> 仅由小写英文字符组成。</li>
</ul>

## 标签

 - String 
 - Dynamic Programming 

## 代码

```golang
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
```