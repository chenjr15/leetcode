# 剑指 Offer II 018.有效的回文

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/XltzEq/](https://leetcode-cn.com/problems/XltzEq/)

## 题目描述

<p>给定一个字符串 <code>s</code> ，验证 <code>s</code>&nbsp;是否是&nbsp;<strong>回文串&nbsp;</strong>，只考虑字母和数字字符，可以忽略字母的大小写。</p>

<p>本题中，将空字符串定义为有效的&nbsp;<strong>回文串&nbsp;</strong>。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入: </strong>s =<strong> </strong>&quot;A man, a plan, a canal: Panama&quot;
<strong>输出:</strong> true
<strong>解释：</strong>&quot;amanaplanacanalpanama&quot; 是回文串</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> s = &quot;race a car&quot;
<strong>输出:</strong> false
解释：&quot;raceacar&quot; 不是回文串</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 2 * 10<sup>5</sup></code></li>
	<li>字符串 <code>s</code> 由 ASCII 字符组成</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 125&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/valid-palindrome/">https://leetcode-cn.com/problems/valid-palindrome/</a></p>

## 标签

 - Two Pointers 
 - String 

## 代码

```golang
func isPalindrome(s string) bool {
    left:=0
    right:=len(s)-1
    var rchar,lchar byte 
    for left<right{
        lchar = s[left]
        // 跳过符号
        for left<right && !isLetterNum(lchar) {
            // fmt.Printf("Passing: %c at %d %v\n",rune(lchar),left,isLetterNum(lchar))
            left++
            lchar = s[left]
        }
        rchar = s[right]
        for left<right && !isLetterNum(rchar) {
            // fmt.Printf("Passing: %c at %d %v\n",rune(rchar),right,isLetterNum(rchar))
            right--
            rchar = s[right]
        }
        rchar = toLowerCase(rchar)
        lchar = toLowerCase(lchar)
        // fmt.Printf("%d,%d,%c,%c\n",left,right,lchar,rchar)
        if rchar != lchar{
            return false
        }
        right--
        left++

    }
    return true

}

func toLowerCase(c byte) byte{
    if  (c>='A' && c <='Z'){
        c += 'a'-'A'
    }
    return c

}

func isLetterNum(c byte)bool{
    if  (c>='A' && c <='Z'){
        return true
    }
    if (c>='a' &&  c <='z'){
        return true
    }
    if (c>='0' &&  c <='9'){
        return true
    }
    return false

}
```