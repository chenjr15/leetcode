# 168.Excel表列名称

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/excel-sheet-column-title/](https://leetcode-cn.com/problems/excel-sheet-column-title/)

## 题目描述

<p>给你一个整数 <code>columnNumber</code> ，返回它在 Excel 表中相对应的列名称。</p>

<p>例如：</p>

<pre>
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28 
...
</pre>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>columnNumber = 1
<strong>输出：</strong>"A"
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>columnNumber = 28
<strong>输出：</strong>"AB"
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>columnNumber = 701
<strong>输出：</strong>"ZY"
</pre>

<p><strong>示例 4：</strong></p>

<pre>
<strong>输入：</strong>columnNumber = 2147483647
<strong>输出：</strong>"FXSHRXW"
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= columnNumber <= 2<sup>31</sup> - 1</code></li>
</ul>

## 标签

 - Math 
 - String 

## 代码

```golang
func convertToTitle(columnNumber int) string {
    r:=0
    runes:=[]rune{}
    
    // columnNumber--
    for columnNumber>0{
        columnNumber--
        r =  columnNumber%26
        columnNumber/=26
        runes  = append(runes,rune(r+'A'))
    }
    // runes  = append(runes,rune(columnNumber+'A'))
    // fmt.Println(runes)
    // reverse
    for i,j:=0,len(runes)-1;i<j;i,j=i+1,j-1{
        runes[i],runes[j] = runes[j],runes[i]
    }
    return string(runes)
}
```