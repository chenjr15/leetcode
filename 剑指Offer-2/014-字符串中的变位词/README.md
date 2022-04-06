# 剑指 Offer II 014.字符串中的变位词

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/MPnaiL/](https://leetcode-cn.com/problems/MPnaiL/)

## 题目描述

<p>给定两个字符串&nbsp;<code>s1</code>&nbsp;和&nbsp;<code>s2</code>，写一个函数来判断 <code>s2</code> 是否包含 <code>s1</code><strong>&nbsp;</strong>的某个变位词。</p>

<p>换句话说，第一个字符串的排列之一是第二个字符串的 <strong>子串</strong> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入: </strong>s1 = &quot;ab&quot; s2 = &quot;eidbaooo&quot;
<strong>输出: </strong>True
<strong>解释:</strong> s2 包含 s1 的排列之一 (&quot;ba&quot;).
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入: </strong>s1= &quot;ab&quot; s2 = &quot;eidboaoo&quot;
<strong>输出:</strong> False
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s1.length, s2.length &lt;= 10<sup>4</sup></code></li>
	<li><code>s1</code> 和 <code>s2</code> 仅包含小写字母</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 567&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/permutation-in-string/">https://leetcode-cn.com/problems/permutation-in-string/</a></p>

## 标签

 - Hash Table 
 - Two Pointers 
 - String 
 - Sliding Window 

## 代码

```golang
func checkInclusion(s1 string, s2 string) bool {
    if len(s1)>len(s2){
        return false
    }
    // s1len := len(s1)
    // s2len := len(s2)
    // 字符串滚动哈希
    s1hash := make([]int,26)
    s2hash := make([]int,26)
    // 计算s1的哈希
    for i :=range []byte(s1){
        s1hash[s1[i] - 'a']++
        s2hash[s2[i] - 'a']++
    }
    if arrayEquals(s1hash,s2hash){
        return true
    }
    difflen := len(s2)-len(s1)+1
    
    for i :=1;i<difflen;i++ {
        // 加头
        s2hash[s2[i+len(s1)-1]-'a']++
        // 去尾
        s2hash[s2[i-1]-'a']--
        if arrayEquals(s1hash,s2hash){
            return true
        }
    }
    return false

}

func arrayEquals(a ,b []int)( equals bool){
    equals=true
    for i:= range a{
        if a[i] != b[i]{
            equals =false
            break
        }
    }
    return equals

}
```