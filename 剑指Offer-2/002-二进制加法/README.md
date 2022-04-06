# 剑指 Offer II 002.二进制加法

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/JFETK5/](https://leetcode-cn.com/problems/JFETK5/)

## 题目描述

<p>给定两个 01 字符串&nbsp;<code>a</code>&nbsp;和&nbsp;<code>b</code>&nbsp;，请计算它们的和，并以二进制字符串的形式输出。</p>

<p>输入为 <strong>非空 </strong>字符串且只包含数字&nbsp;<code>1</code>&nbsp;和&nbsp;<code>0</code>。</p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre>
<strong>输入:</strong> a = &quot;11&quot;, b = &quot;10&quot;
<strong>输出:</strong> &quot;101&quot;</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre>
<strong>输入:</strong> a = &quot;1010&quot;, b = &quot;1011&quot;
<strong>输出:</strong> &quot;10101&quot;</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>每个字符串仅由字符 <code>&#39;0&#39;</code> 或 <code>&#39;1&#39;</code> 组成。</li>
	<li><code>1 &lt;= a.length, b.length &lt;= 10^4</code></li>
	<li>字符串如果不是 <code>&quot;0&quot;</code> ，就都不含前导零。</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 67&nbsp;题相同：<a href="https://leetcode-cn.com/problems/add-binary/">https://leetcode-cn.com/problems/add-binary/</a></p>

## 标签

 - Bit Manipulation 
 - Math 
 - String 
 - Simulation 

## 代码

```golang
func addBinary(a string, b string) string {
    lenA := len(a)
    lenB := len(b)
    // 二进制加法
    maxLen := lenA
    if lenB>maxLen{
        maxLen = lenB
    }
    result:= make([]byte,maxLen,maxLen+1)
    var carry byte
    i:=0
    for ;i<maxLen;i++{
        if i <lenA{
            carry+= a[lenA - i -1] - '0'
        }
        if i <lenB{
            carry+= b[lenB - i -1] - '0'
        }
        if carry & 1 != 0{
            result[i] = '1'
        }else{
            result[i] = '0'
        }
        carry>>=1
    }
    if carry>0{
        result = append(result,'1')
    }
    // reverse
    l:=len(result)
    for i,j:=0,l-1;i<j;{
        result[i],result[j] = result[j],result[i]
        i++
        j--
    }
    return string(result)

}
```