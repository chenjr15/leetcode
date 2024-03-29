# 剑指 Offer 56 - I.数组中数字出现的次数

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/)

## 题目描述

<p>一个整型数组 <code>nums</code> 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [4,1,4,6]
<strong>输出：</strong>[1,6] 或 [6,1]
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [1,2,10,4,1,4,3,3]
<strong>输出：</strong>[2,10] 或 [10,2]</pre>

<p>&nbsp;</p>

<p><strong>限制：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 10000</code></li>
</ul>

<p>&nbsp;</p>

## 标签

 - Bit Manipulation 
 - Array 

## 代码

```golang
func singleNumbers(nums []int) []int {
    var sum,x1,x2 int
    // 求出sum=x1^x2
    for _,n:=range nums{
        sum^= n
    }
    // 求最低有效位
    lsb:=sum & -sum
    for _,n:=range nums{
        // 按lsb分类异或
        if n & lsb !=0{
            x1^=n
        }else {
            x2^=n
        }
    }
    return []int{x1,x2}

}
```