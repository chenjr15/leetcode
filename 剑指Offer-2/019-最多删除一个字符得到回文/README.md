# 剑指 Offer II 019.最多删除一个字符得到回文

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/RQku0D/](https://leetcode-cn.com/problems/RQku0D/)

## 题目描述

<p>给定一个非空字符串&nbsp;<code>s</code>，请判断如果&nbsp;<strong>最多 </strong>从字符串中删除一个字符能否得到一个回文字符串。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> s = &quot;aba&quot;
<strong>输出:</strong> true
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> s = &quot;abca&quot;
<strong>输出:</strong> true
<strong>解释:</strong> 可以删除 &quot;c&quot; 字符 或者 &quot;b&quot; 字符
</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入:</strong> s = &quot;abc&quot;
<strong>输出:</strong> false</pre>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s</code> 由小写英文字母组成</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 680&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/valid-palindrome-ii/">https://leetcode-cn.com/problems/valid-palindrome-ii/</a></p>

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