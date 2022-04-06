# 537.复数乘法

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/complex-number-multiplication/](https://leetcode-cn.com/problems/complex-number-multiplication/)

## 题目描述

<p><a href="https://baike.baidu.com/item/%E5%A4%8D%E6%95%B0/254365?fr=aladdin" target="_blank">复数</a> 可以用字符串表示，遵循 <code>"<strong>实部</strong>+<strong>虚部</strong>i"</code> 的形式，并满足下述条件：</p>

<ul>
	<li><code>实部</code> 是一个整数，取值范围是 <code>[-100, 100]</code></li>
	<li><code>虚部</code> 也是一个整数，取值范围是 <code>[-100, 100]</code></li>
	<li><code>i<sup>2</sup> == -1</code></li>
</ul>

<p>给你两个字符串表示的复数 <code>num1</code> 和 <code>num2</code> ，请你遵循复数表示形式，返回表示它们乘积的字符串。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>num1 = "1+1i", num2 = "1+1i"
<strong>输出：</strong>"0+2i"
<strong>解释：</strong>(1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i ，你需要将它转换为 0+2i 的形式。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>num1 = "1+-1i", num2 = "1+-1i"
<strong>输出：</strong>"0+-2i"
<strong>解释：</strong>(1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i ，你需要将它转换为 0+-2i 的形式。 
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>num1</code> 和 <code>num2</code> 都是有效的复数表示。</li>
</ul>

## 标签

 - Math 
 - String 
 - Simulation 

## 代码

```golang
func complexNumberMultiply(num1 string, num2 string) string {

    // a+bi , c+di
    var a,b,c,d int
    // e+fi = (a+bi)* (c+di) = ac + cbi + adi - bd
    var e,f int

    snum := strings.Split(num1[:len(num1)-1],"+")
    a,err :=strconv.Atoi(snum[0])
    if err != nil{
        fmt.Println(err)
    }
    b,err=strconv.Atoi(snum[1])
    if err != nil{
        fmt.Println(err)
    }
    snum = strings.Split(num2[:len(num2)-1],"+")
    c,err=strconv.Atoi(snum[0])
    if err != nil{
        fmt.Println(err)
    }
    d,err=strconv.Atoi(snum[1])
    if err != nil{
        fmt.Println(err)
    }
    e = a*c -b*d
    f = c*b + a*d
    return fmt.Sprintf("%d+%di",e,f)

}
```