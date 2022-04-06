# 剑指 Offer II 020.回文子字符串的个数

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/a7VOhD/](https://leetcode-cn.com/problems/a7VOhD/)

## 题目描述

<p>给定一个字符串 <code>s</code> ，请计算这个字符串中有多少个回文子字符串。</p>

<p>具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "abc"
<strong>输出：</strong>3
<strong>解释：</strong>三个回文子串: "a", "b", "c"
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>s =<strong> </strong>"aaa"
<strong>输出：</strong>6
<strong>解释：</strong>6个回文子串: "a", "a", "a", "aa", "aa", "aaa"</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 1000</code></li>
	<li><code>s</code> 由小写英文字母组成</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 647 题相同：<a href="https://leetcode-cn.com/problems/palindromic-substrings/">https://leetcode-cn.com/problems/palindromic-substrings/</a>&nbsp;</p>

## 标签

 - String 
 - Dynamic Programming 

## 代码

```golang
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
```