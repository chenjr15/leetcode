# 260.只出现一次的数字 III

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/single-number-iii/](https://leetcode-cn.com/problems/single-number-iii/)

## 题目描述

<p>给定一个整数数组 <code>nums</code>，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 <strong>任意顺序</strong> 返回答案。</p>

<p> </p>

<p><strong>进阶：</strong>你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,2,1,3,2,5]
<strong>输出：</strong>[3,5]
<strong>解释：</strong>[5, 3] 也是有效的答案。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [-1,0]
<strong>输出：</strong>[-1,0]
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums = [0,1]
<strong>输出：</strong>[1,0]
</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 <= nums.length <= 3 * 10<sup>4</sup></code></li>
	<li><code>-2<sup>31</sup> <= nums[i] <= 2<sup>31</sup> - 1</code></li>
	<li>除两个只出现一次的整数外，<code>nums</code> 中的其他数字都出现两次</li>
</ul>

## 标签

 - Bit Manipulation 
 - Array 

## 代码

```golang
func singleNumber(nums []int) []int {
    sum:= 0
    x1 := 0
    x2 := 0
    // 根据答案的思路
    // 将所有的元素异或一次以后，因为对同一个数异或两次能够消除这个数，所以最后异或的结果已经是x1^x2 即只剩下x1和x2
    for _,v :=range nums{
        sum^=v
    }
    // 接下去是想办法获得x1 和x2
    // 首先x1 不等于x2， 因此x1 和x2 异或必定不等于0，即sum不为0
    // 用lowbit求出sum最低位的1的位置为k，那么这个最低位的1。那么在x1 和x2中对应的第k位肯定相等。
    // 为此我们将该位为0 和为1数分开，x1 和x2 肯定分别在两个类中。同时，其他元素因为都有两个，这两个一定在同一个类中
    // 因此只要对这两个类分别异或即可得到对应的元素
    lowbit:= sum & -sum
    for _,v:=range nums{
        if (v&lowbit) == 0{
            x1 ^= v
        }else{
            x2 ^= v
        }
    }
    return []int{x1,x2}

}
```