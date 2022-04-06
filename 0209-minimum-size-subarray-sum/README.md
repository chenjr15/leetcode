# 209.长度最小的子数组

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/minimum-size-subarray-sum/](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)

## 题目描述

<p>给定一个含有 <code>n</code><strong> </strong>个正整数的数组和一个正整数 <code>target</code><strong> 。</strong></p>

<p>找出该数组中满足其和<strong> </strong><code>≥ target</code><strong> </strong>的长度最小的 <strong>连续子数组</strong> <code>[nums<sub>l</sub>, nums<sub>l+1</sub>, ..., nums<sub>r-1</sub>, nums<sub>r</sub>]</code> ，并返回其长度<strong>。</strong>如果不存在符合条件的子数组，返回 <code>0</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>target = 7, nums = [2,3,1,2,4,3]
<strong>输出：</strong>2
<strong>解释：</strong>子数组 <code>[4,3]</code> 是该条件下的长度最小的子数组。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>target = 4, nums = [1,4,4]
<strong>输出：</strong>1
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>target = 11, nums = [1,1,1,1,1,1,1,1]
<strong>输出：</strong>0
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= target <= 10<sup>9</sup></code></li>
	<li><code>1 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>5</sup></code></li>
</ul>

<p> </p>

<p><strong>进阶：</strong></p>

<ul>
	<li>如果你已经实现<em> </em><code>O(n)</code> 时间复杂度的解法, 请尝试设计一个 <code>O(n log(n))</code> 时间复杂度的解法。</li>
</ul>

## 标签

 - Array 
 - Binary Search 
 - Prefix Sum 
 - Sliding Window 

## 代码

```golang
func minSubArrayLen(target int, nums []int) int {
    minlen:= 0
    nlen:=len(nums)
    // 尝试滑动窗口
    // 首先对范围有一个明确的定义，子数组的范围就是[left,right],那么sum=nums[left,...,right],
    // curlen = right-left +1, left<=right < nlen 
    var left,right,sum,curlen int
    sum = nums[0]
    curlen=1
    for left <= right {
        // 往右扩直到范围内的数组和大于target
        for right<nlen-1 && sum < target{
            right++
            sum += nums[right]
        }
        
        for left <= right && sum>= target{
            curlen= right-left+1 
            if minlen == 0 || curlen<minlen{
                minlen = curlen
            }
            // 往左缩
            sum -= nums[left]
            left++
        }
        // 跳出的关键，右边扩张到最后一个，且左边已经收缩完了，就可以退出了
        if right==nlen-1{
            break
        }
        
    }
    return minlen

}
```