# 137.只出现一次的数字 II

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/single-number-ii/](https://leetcode-cn.com/problems/single-number-ii/)

## 题目描述

<p>给你一个整数数组 <code>nums</code> ，除某个元素仅出现 <strong>一次</strong> 外，其余每个元素都恰出现 <strong>三次 。</strong>请你找出并返回那个只出现了一次的元素。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [2,2,3,2]
<strong>输出：</strong>3
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [0,1,0,1,0,1,99]
<strong>输出：</strong>99
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 3 * 10<sup>4</sup></code></li>
	<li><code>-2<sup>31</sup> <= nums[i] <= 2<sup>31</sup> - 1</code></li>
	<li><code>nums</code> 中，除某个元素仅出现 <strong>一次</strong> 外，其余每个元素都恰出现 <strong>三次</strong></li>
</ul>

<p> </p>

<p><strong>进阶：</strong>你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？</p>

## 标签

 - Bit Manipulation 
 - Array 

## 代码

```golang
func singleNumber(nums []int) int {
    result := 0
    sum:=0
    for i:=0;i<32;i++{
        sum=0
        for _,v:=range nums{
            sum+= (v>>i)&1
        }
        sum%=3
        result |= (sum&1)<<i

    }

    return int(int32(result))
 
}
```