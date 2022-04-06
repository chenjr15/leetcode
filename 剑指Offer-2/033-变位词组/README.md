# 剑指 Offer II 033.变位词组

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/sfvd7V/](https://leetcode-cn.com/problems/sfvd7V/)

## 题目描述

<p>给定一个字符串数组 <code>strs</code> ，将&nbsp;<strong>变位词&nbsp;</strong>组合在一起。 可以按任意顺序返回结果列表。</p>

<p><strong>注意：</strong>若两个字符串中每个字符出现的次数都相同，则称它们互为变位词。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> strs = <code>[&quot;eat&quot;, &quot;tea&quot;, &quot;tan&quot;, &quot;ate&quot;, &quot;nat&quot;, &quot;bat&quot;]</code>
<strong>输出: </strong>[[&quot;bat&quot;],[&quot;nat&quot;,&quot;tan&quot;],[&quot;ate&quot;,&quot;eat&quot;,&quot;tea&quot;]]</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> strs = <code>[&quot;&quot;]</code>
<strong>输出: </strong>[[&quot;&quot;]]
</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入:</strong> strs = <code>[&quot;a&quot;]</code>
<strong>输出: </strong>[[&quot;a&quot;]]</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= strs.length &lt;= 10<sup>4</sup></code></li>
	<li><code>0 &lt;= strs[i].length &lt;= 100</code></li>
	<li><code>strs[i]</code>&nbsp;仅包含小写字母</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 49&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/group-anagrams/">https://leetcode-cn.com/problems/group-anagrams/</a></p>

## 标签

 - Hash Table 
 - String 
 - Sorting 

## 代码

```golang
func groupAnagrams(strs []string) [][]string {
    // 考虑用字符哈希来逐个比较，
    // 逐个比较的过程中可以给他排序按长度排序，然后按哈希排序
    // 实在不行把每个字符串排序
    // -------------
    // 考虑使用字符串压缩方法，来做哈希表的key
    strHash:= map[[26]int][]string {}

    for _,str:=range strs{
        // 计算压缩哈希
        cntKey:=[26]int{}
        for i:=range str{
            cntKey[str[i]-'a'] ++
        }
        list:= strHash[cntKey]
        if list==nil{
            list=[]string{str}
        }else{
            list = append(list,str)
        }
        strHash[cntKey] = list
    }
    reslut := [][]string{}
    for _,str:= range strHash{
        reslut=append(reslut,str)
    }
    return reslut
}


```