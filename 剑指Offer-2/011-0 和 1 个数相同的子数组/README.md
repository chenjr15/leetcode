# 剑指 Offer II 011.0 和 1 个数相同的子数组

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/A1NYOS/](https://leetcode-cn.com/problems/A1NYOS/)

## 题目描述

<p>给定一个二进制数组 <code>nums</code> , 找到含有相同数量的 <code>0</code> 和 <code>1</code> 的最长连续子数组，并返回该子数组的长度。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> nums = [0,1]
<strong>输出:</strong> 2
<strong>说明:</strong> [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> nums = [0,1,0]
<strong>输出:</strong> 2
<strong>说明:</strong> [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>nums[i]</code> 不是 <code>0</code> 就是 <code>1</code></li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 525&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/contiguous-array/">https://leetcode-cn.com/problems/contiguous-array/</a></p>

## 标签

 - Array 
 - Hash Table 
 - Prefix Sum 

## 代码

```golang
func findMaxLength(nums []int) int {

    // 指定区间1 的个数可以用前缀和来算，还是和上一题类似
    // 如果指定区间内0和1个数相等，则区间长度一定是偶数，并且区间的长度=两倍的区间和
    // 用-1 替换0就可以直接使01相等的数组的和变为0
    psum:=0
    psumPos:= make(map[int]int)
    maxLen:= 0
    for i := range nums{
        // i到 n-1 ,因此需要遍历的是位置2的情况
        // 位置1 在这里psum表达的是[0,i-1]的前缀和，长度是i
        if nums[i]==0{
            psum--

        }else{
            psum++
        }
        // 如果前缀和是0，那么0-当前下标的这段闭区间就是01个数相等，同时以i为结尾的不会有更长的连续子数组了
        if psum==0{
            if i+1>maxLen{
                maxLen=i+1
            }
        }else{
            // 否则去寻找前面和当前前缀和相等的位置
            if start,ok := psumPos[psum];ok{
                // 存在，则用这个去减对应的位置
                // 存在一定是最小的下标，即通过这个下标可以得到最长的子数组
                if i-start>maxLen{
                    maxLen=i-start
                }
            }else{
                // 不存在，将当前的前缀和加入哈希表
                psumPos[psum] = i 
            }
        }
        
    }
    return maxLen
}
```