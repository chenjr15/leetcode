# 67.二进制求和

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/add-binary/](https://leetcode-cn.com/problems/add-binary/)

## 题目描述

<p>给你两个二进制字符串，返回它们的和（用二进制表示）。</p>

<p>输入为 <strong>非空 </strong>字符串且只包含数字&nbsp;<code>1</code>&nbsp;和&nbsp;<code>0</code>。</p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> a = &quot;11&quot;, b = &quot;1&quot;
<strong>输出:</strong> &quot;100&quot;</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> a = &quot;1010&quot;, b = &quot;1011&quot;
<strong>输出:</strong> &quot;10101&quot;</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>每个字符串仅由字符 <code>&#39;0&#39;</code> 或 <code>&#39;1&#39;</code> 组成。</li>
	<li><code>1 &lt;= a.length, b.length &lt;= 10^4</code></li>
	<li>字符串如果不是 <code>&quot;0&quot;</code> ，就都不含前导零。</li>
</ul>

## 标签

 - Bit Manipulation 
 - Math 
 - String 
 - Simulation 

## 代码

```java
class Solution {
    public String addBinary(String a, String b) {
        int pa = a.length() - 1;
        int pb = b.length() - 1;
        int carry = 0;
        int sum = 0 ;
        StringBuilder  sb = new StringBuilder();
        while(pa>-1 && pb >-1 ){
            sum = a.charAt(pa) - '0' + b.charAt(pb) -'0' + carry;
            carry = sum >>1;
            sum %=2;
            sb.append((char)('0'+sum));
            // System.out.println("" + sum%2 + " "+carry+" sb:"+sb.toString());
            
            pa--;
            pb--;
        }
        String remain = a;
        int pr = pa; 
        if( pa == -1 ){
            pr=pb;
            remain = b;
        }
        
        for(;pr>-1;pr--){
            sum = remain.charAt(pr) - '0' + carry;
            // System.out.println("" + sum%2 + " "+carry+" sb:"+sb.toString());

            sb.append((char)('0'+(sum %2)));
            carry = sum >>1;
        }
        if(carry>0){
            sb.append('1');
        }
        
        return sb.reverse().toString();
    }
}
```