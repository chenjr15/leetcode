# 剑指 Offer II 015.字符串中的所有变位词

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/VabMRr/](https://leetcode-cn.com/problems/VabMRr/)

## 题目描述

<p>给定两个字符串&nbsp;<code>s</code>&nbsp;和<b>&nbsp;</b><code>p</code>，找到&nbsp;<code>s</code><strong>&nbsp;</strong>中所有 <code>p</code> 的&nbsp;<strong>变位词&nbsp;</strong>的子串，返回这些子串的起始索引。不考虑答案输出的顺序。</p>

<p><strong>变位词 </strong>指字母相同，但排列不同的字符串。</p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre>
<strong>输入: </strong>s = &quot;cbaebabacd&quot;, p = &quot;abc&quot;
<strong>输出: </strong>[0,6]
<strong>解释:</strong>
起始索引等于 0 的子串是 &quot;cba&quot;, 它是 &quot;abc&quot; 的变位词。
起始索引等于 6 的子串是 &quot;bac&quot;, 它是 &quot;abc&quot; 的变位词。
</pre>

<p><strong>&nbsp;示例 2:</strong></p>

<pre>
<strong>输入: </strong>s = &quot;abab&quot;, p = &quot;ab&quot;
<strong>输出: </strong>[0,1,2]
<strong>解释:</strong>
起始索引等于 0 的子串是 &quot;ab&quot;, 它是 &quot;ab&quot; 的变位词。
起始索引等于 1 的子串是 &quot;ba&quot;, 它是 &quot;ab&quot; 的变位词。
起始索引等于 2 的子串是 &quot;ab&quot;, 它是 &quot;ab&quot; 的变位词。
</pre>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 &lt;= s.length, p.length &lt;= 3 * 10<sup>4</sup></code></li>
	<li><code>s</code>&nbsp;和 <code>p</code> 仅包含小写字母</li>
</ul>

<p>&nbsp;</p>

<p>注意：本题与主站 438&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/" style="background-color: rgb(255, 255, 255);">https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/</a></p>

## 标签

 - Hash Table 
 - String 
 - Sliding Window 

## 代码

```golang
func findAnagrams(s2 string, s1 string) []int {
    result:= make([]int,0)
    if len(s1)>len(s2){
        return result
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
        result=append(result,0)
    }
    difflen := len(s2)-len(s1)+1
    
    for i :=1;i<difflen;i++ {
        // 加头
        s2hash[s2[i+len(s1)-1]-'a']++
        // 去尾
        s2hash[s2[i-1]-'a']--
        if arrayEquals(s1hash,s2hash){
            result=append(result,i)
        }
    }
    return result

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