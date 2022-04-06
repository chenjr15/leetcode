# 680.验证回文字符串 Ⅱ

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/valid-palindrome-ii/](https://leetcode-cn.com/problems/valid-palindrome-ii/)

## 题目描述

<p>给定一个非空字符串 <code>s</code>，<strong>最多</strong>删除一个字符。判断是否能成为回文字符串。</p>

<p> </p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> s = "aba"
<strong>输出:</strong> true
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> s = "abca"
<strong>输出:</strong> true
<strong>解释:</strong> 你可以删除c字符。
</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入:</strong> s = "abc"
<strong>输出:</strong> false</pre>

<p> </p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 <= s.length <= 10<sup>5</sup></code></li>
	<li><code>s</code> 由小写英文字母组成</li>
</ul>

## 标签

 - Greedy 
 - Two Pointers 
 - String 

## 代码

```golang
func validPalindrome(s string) bool {
    // 题目说只改一个字符，因此我们在碰到不是回文的字符的情况下，
    // 尝试将左边或者右边删掉，看是否能够形成回文川，也就是只要试两次
    left:=0
    right:=len(s)-1
    for left<right{

        // fmt.Printf("%d,%d,%c,%c\n",left,right,lchar,rchar)
        if s[left] != s[right]{
            // 尝试判断左边或者右边删掉的情况
            // 删掉左边
            if !isPalindrome(s,left+1,right){
                return isPalindrome(s,left,right-1)
            }else {
                return true
            }
        }
        right--
        left++
    }
    return true

}

func isPalindrome(s string, left ,right int ) bool{
    for left<right{
        // fmt.Printf("%d,%d,%c,%c\n",left,right,lchar,rchar)
        if s[left] != s[right]{
            // 尝试判断左边或者右边删掉的情况
            // 删掉左边
            return false 
        }
        right--
        left++
    }
    return true
}
```