# 剑指 Offer II 005.单词长度的最大乘积

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/aseY1I/](https://leetcode-cn.com/problems/aseY1I/)

## 题目描述

<p>给定一个字符串数组&nbsp;<code>words</code>，请计算当两个字符串 <code>words[i]</code> 和 <code>words[j]</code> 不包含相同字符时，它们长度的乘积的最大值。假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。</p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre>
<strong>输入:</strong> words = <code>[&quot;abcw&quot;,&quot;baz&quot;,&quot;foo&quot;,&quot;bar&quot;,&quot;fxyz&quot;,&quot;abcdef&quot;]</code>
<strong>输出: </strong><code>16 
<strong>解释:</strong> 这两个单词为<strong> </strong></code><code>&quot;abcw&quot;, &quot;fxyz&quot;</code>。它们不包含相同字符，且长度的乘积最大。</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> words = <code>[&quot;a&quot;,&quot;ab&quot;,&quot;abc&quot;,&quot;d&quot;,&quot;cd&quot;,&quot;bcd&quot;,&quot;abcd&quot;]</code>
<strong>输出: </strong><code>4 
<strong>解释: </strong></code>这两个单词为 <code>&quot;ab&quot;, &quot;cd&quot;</code>。</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入:</strong> words = <code>[&quot;a&quot;,&quot;aa&quot;,&quot;aaa&quot;,&quot;aaaa&quot;]</code>
<strong>输出: </strong><code>0 
<strong>解释: </strong>不存在这样的两个单词。</code>
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= words.length &lt;= 1000</code></li>
	<li><code>1 &lt;= words[i].length &lt;= 1000</code></li>
	<li><code>words[i]</code>&nbsp;仅包含小写字母</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 318&nbsp;题相同：<a href="https://leetcode-cn.com/problems/maximum-product-of-word-lengths/">https://leetcode-cn.com/problems/maximum-product-of-word-lengths/</a></p>

## 标签

 - Bit Manipulation 
 - Array 
 - String 

## 代码

```golang
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
```