# 剑指 Offer II 004.只出现一次的数字

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/WGki4K/](https://leetcode-cn.com/problems/WGki4K/)

## 题目描述

<p>给你一个整数数组&nbsp;<code>nums</code> ，除某个元素仅出现 <strong>一次</strong> 外，其余每个元素都恰出现 <strong>三次 。</strong>请你找出并返回那个只出现了一次的元素。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [2,2,3,2]
<strong>输出：</strong>3
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [0,1,0,1,0,1,100]
<strong>输出：</strong>100
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 3 * 10<sup>4</sup></code></li>
	<li><code>-2<sup>31</sup> &lt;= nums[i] &lt;= 2<sup>31</sup> - 1</code></li>
	<li><code>nums</code> 中，除某个元素仅出现 <strong>一次</strong> 外，其余每个元素都恰出现 <strong>三次</strong></li>
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？</p>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 137&nbsp;题相同：<a href="https://leetcode-cn.com/problems/single-number-ii/">https://leetcode-cn.com/problems/single-number-ii/</a></p>

## 标签

 - Bit Manipulation 
 - Array 

## 代码

```golang
func singleNumber(nums []int) int {
    result :=0
    // 按二进制逐位计算
    // 对于每一位来说，如果把所有数字对应的1加起来的话得到和为sum。
    // 因为除了目标元素以为，其他元素都出现三次，如果那个元素对应的二进制为1 那么一定使sum加3 如果是0则无影响。
    // 同样的，目标元素会使sum+1。其他元素对sum的影响可以用过对3区域来去除，只剩下目标元素带来影响的1 或者0
    sum:=0
    for i:=0;i<32;i++{
        sum =0 
        for _,v:=range nums{
            sum+=((v>>i)&1)
        }
        // 去除该元素的影响
        result |= ((sum%3)<<i  )
    }
    return int(int32(result))
    
}
```