# 剑指 Offer II 009.乘积小于 K 的子数组

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/ZVAVXX/](https://leetcode-cn.com/problems/ZVAVXX/)

## 题目描述

<p>给定一个正整数数组&nbsp;<code>nums</code>和整数 <code>k</code>&nbsp;，请找出该数组内乘积小于&nbsp;<code>k</code>&nbsp;的连续的子数组的个数。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> nums = [10,5,2,6], k = 100
<strong>输出:</strong> 8
<strong>解释:</strong> 8 个乘积小于 100 的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> nums = [1,2,3], k = 0
<strong>输出:</strong> 0</pre>

<p>&nbsp;</p>

<p><strong>提示:&nbsp;</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 3 * 10<sup>4</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 1000</code></li>
	<li><code>0 &lt;= k &lt;= 10<sup>6</sup></code></li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 713&nbsp;题相同：<a href="https://leetcode-cn.com/problems/subarray-product-less-than-k/">https://leetcode-cn.com/problems/subarray-product-less-than-k/</a>&nbsp;</p>

## 标签

 - Array 
 - Sliding Window 

## 代码

```golang
func numSubarrayProductLessThanK(nums []int, target int) int {
    // 应该是滑动窗口,拿上一题过来套

    arrCnt:= 0
    nlen:=len(nums)
    // 尝试滑动窗口
    // 首先对范围有一个明确的定义，子数组的范围就是[left,right),那么sum=nums[left,...,right-1],
    // curlen = right-left, left<right <= nlen 
    var left,right,product int
    right=1
    product = nums[0]

    for left < right {
        // 往右扩直到范围内的乘积大于等于target
        for right<nlen && product < target{
            product *= nums[right]
            right++
        }
        // right 是第一个不满足条件的位置，因此
        // fmt.Println(left,right,product)
        // 得到以left开头的，right-1结尾的子数组乘积满足条件
        arrCnt+= right-left-1
        if product < target{
            arrCnt++
        }
        if left == nlen{
            break
        }
        product/=nums[left]
        left++
        // for left < right && sum>= target{
        //     curlen= right-left+1 
        //     if minlen == 0 || curlen<minlen{
        //         minlen = curlen
        //     }
        //     // 往左缩
        //     sum -= nums[left]
        //     left++
        // }
        // // 跳出的关键，右边扩张到最后一个，且左边已经收缩完了，就可以退出了
        // if right==nlen-1{
        //     break
        // }
        
    }
    return arrCnt

}
```