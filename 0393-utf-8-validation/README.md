# 393.UTF-8 编码验证

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/utf-8-validation/](https://leetcode-cn.com/problems/utf-8-validation/)

## 题目描述

<p>给定一个表示数据的整数数组&nbsp;<code>data</code>&nbsp;，返回它是否为有效的 <strong>UTF-8</strong> 编码。</p>

<p><strong>UTF-8</strong> 中的一个字符可能的长度为 <strong>1 到 4 字节</strong>，遵循以下的规则：</p>

<ol>
	<li>对于 <strong>1 字节</strong>&nbsp;的字符，字节的第一位设为 0 ，后面 7 位为这个符号的 unicode 码。</li>
	<li>对于 <strong>n 字节</strong>&nbsp;的字符 (n &gt; 1)，第一个字节的前 n 位都设为1，第 n+1 位设为 0 ，后面字节的前两位一律设为 10 。剩下的没有提及的二进制位，全部为这个符号的 unicode 码。</li>
</ol>

<p>这是 UTF-8 编码的工作方式：</p>

<pre>
<code>   Char. number range  |        UTF-8 octet sequence
      (hexadecimal)    |              (binary)
   --------------------+---------------------------------------------
   0000 0000-0000 007F | 0xxxxxxx
   0000 0080-0000 07FF | 110xxxxx 10xxxxxx
   0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
   0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
</code></pre>

<p><strong>注意：</strong>输入是整数数组。只有每个整数的 <strong>最低 8 个有效位</strong> 用来存储数据。这意味着每个整数只表示 1 字节的数据。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>data = [197,130,1]
<strong>输出：</strong>true
<strong>解释：</strong>数据表示字节序列:<strong>11000101 10000010 00000001</strong>。
这是有效的 utf-8 编码，为一个 2 字节字符，跟着一个 1 字节字符。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>data = [235,140,4]
<strong>输出：</strong>false
<strong>解释：</strong>数据表示 8 位的序列: <strong>11101011 10001100 00000100</strong>.
前 3 位都是 1 ，第 4 位为 0 表示它是一个 3 字节字符。
下一个字节是开头为 10 的延续字节，这是正确的。
但第二个延续字节不以 10 开头，所以是不符合规则的。
</pre>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 &lt;= data.length &lt;= 2 * 10<sup>4</sup></code></li>
	<li><code>0 &lt;= data[i] &lt;= 255</code></li>
</ul>

## 标签

 - Bit Manipulation 
 - Array 

## 代码

```golang
func validUtf8(data []int) bool {

    dataLenMask:=[]int{
        0b10000000,
        0b11100000,
        0b11110000,
        0b11111000,
    }
    correctData:= []int{
        0b00000000,
        0b11000000,
        0b11100000,
        0b11110000,
    }
    idx:= 0
    runeLen := 0
    for idx<len(data){
        // 这里应该是编码的开头
        // 首先判断这个开头是不是1来决定字节数
        runeLen=0
        for k,mask:=range dataLenMask{
            // fmt.Println(data[idx],mask&data[idx],correctData[k])
            if mask&data[idx] == correctData[k]{
                runeLen=k+1
                break
            }
        }
        // 长度为0 即未匹配上或者长度超过序列长度
        if  runeLen == 0 || ( runeLen+idx)>len(data) {
                // fmt.Println(idx,data[idx],runeLen)
            return false
        }
        // 组
        for k:=1;k<runeLen;k++{
            if data[idx+k] &0b11000000 != 0b10000000{
                // fmt.Println(idx,data[idx+k])
                return false
            }
        }
        idx+=runeLen
    }
    return true

}
```