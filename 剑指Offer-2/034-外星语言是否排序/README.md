# 剑指 Offer II 034.外星语言是否排序

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/lwyVBB/](https://leetcode-cn.com/problems/lwyVBB/)

## 题目描述

<p>某种外星语也使用英文小写字母，但可能顺序 <code>order</code> 不同。字母表的顺序（<code>order</code>）是一些小写字母的排列。</p>

<p>给定一组用外星语书写的单词 <code>words</code>，以及其字母表的顺序 <code>order</code>，只有当给定的单词在这种外星语中按字典序排列时，返回 <code>true</code>；否则，返回 <code>false</code>。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>words = [&quot;hello&quot;,&quot;leetcode&quot;], order = &quot;hlabcdefgijkmnopqrstuvwxyz&quot;
<strong>输出：</strong>true
<strong>解释：</strong>在该语言的字母表中，&#39;h&#39; 位于 &#39;l&#39; 之前，所以单词序列是按字典序排列的。</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>words = [&quot;word&quot;,&quot;world&quot;,&quot;row&quot;], order = &quot;worldabcefghijkmnpqstuvxyz&quot;
<strong>输出：</strong>false
<strong>解释：</strong>在该语言的字母表中，&#39;d&#39; 位于 &#39;l&#39; 之后，那么 words[0] &gt; words[1]，因此单词序列不是按字典序排列的。</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>words = [&quot;apple&quot;,&quot;app&quot;], order = &quot;abcdefghijklmnopqrstuvwxyz&quot;
<strong>输出：</strong>false
<strong>解释：</strong>当前三个字符 &quot;app&quot; 匹配时，第二个字符串相对短一些，然后根据词典编纂规则 &quot;apple&quot; &gt; &quot;app&quot;，因为 &#39;l&#39; &gt; &#39;&empty;&#39;，其中 &#39;&empty;&#39; 是空白字符，定义为比任何其他字符都小（<a href="https://baike.baidu.com/item/%E5%AD%97%E5%85%B8%E5%BA%8F" target="_blank">更多信息</a>）。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= words.length &lt;= 100</code></li>
	<li><code>1 &lt;= words[i].length &lt;= 20</code></li>
	<li><code>order.length == 26</code></li>
	<li>在&nbsp;<code>words[i]</code>&nbsp;和&nbsp;<code>order</code>&nbsp;中的所有字符都是英文小写字母。</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 953&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/verifying-an-alien-dictionary/">https://leetcode-cn.com/problems/verifying-an-alien-dictionary/</a></p>

## 标签

 - Array 
 - Hash Table 
 - String 

## 代码

```golang
func isAlienSorted(words []string, order string) bool {
    if len(words) ==0 {
        return true
    }
    // 将顺序重映射
    orderMap := make([]int,26)
    for i:=range order{
        orderMap[order[i]-'a'] = i
    }
    // 比较字典序, 对于相同的字符跳过，知道比到一个不同的字符
    lenWords := len(words)-1
    for i:=0;i<lenWords;i++{
        
        // fmt.Println(i,words[i],words[i+1],strcmp(words[i],words[i+1],orderMap))
       if  !strcmp(words[i],words[i+1],orderMap) {
           fmt.Println(words[i],words[i+1])
           return false
       }
    }
    return true
}

func strcmp(s1,s2 string, order []int) bool{
    // len(order) == 26
    var p int
    minlen:=len(s1)
    if len(s2)<minlen{
        minlen = len(s2)
    }
    for p=0;p<minlen && s1[p] == s2[p];p++{
        
    }
    // 开始碰到不一样的字符,或者是字符串结尾
    if p==minlen{
        // 走到至少一个字符串的结尾,
        // 字符串1 比字符串2 长,返回false
        fmt.Println(s1,s2, p)

        return len(s1)<=len(s2)
    }
    // 都没有走到结尾, 且不相等，s1的字符比较小则说明是合理的,否则是不合理的
    // fmt.Println(s1,s2, s1[p],s2[p],order[s1[p]-'a'] , order[s2[p]-'a'])
    return order[s1[p]-'a'] < order[s2[p]-'a']
    
}
```