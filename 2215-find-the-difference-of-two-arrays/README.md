# 2215.找出两数组的不同

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/find-the-difference-of-two-arrays/](https://leetcode-cn.com/problems/find-the-difference-of-two-arrays/)

## 题目描述

<p>给你两个下标从 <code>0</code> 开始的整数数组 <code>nums1</code> 和 <code>nums2</code> ，请你返回一个长度为 <code>2</code> 的列表 <code>answer</code> ，其中：</p>

<ul>
	<li><code>answer[0]</code> 是 <code>nums1</code> 中所有<strong> 不 </strong>存在于 <code>nums2</code> 中的 <strong>不同</strong> 整数组成的列表。</li>
	<li><code>answer[1]</code> 是 <code>nums2</code> 中所有<strong> 不 </strong>存在于 <code>nums1</code> 中的 <strong>不同</strong> 整数组成的列表。</li>
</ul>

<p><strong>注意：</strong>列表中的整数可以按 <strong>任意</strong> 顺序返回。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [1,2,3], nums2 = [2,4,6]
<strong>输出：</strong>[[1,3],[4,6]]
<strong>解释：
</strong>对于 nums1 ，nums1[1] = 2 出现在 nums2 中下标 0 处，然而 nums1[0] = 1 和 nums1[2] = 3 没有出现在 nums2 中。因此，answer[0] = [1,3]。
对于 nums2 ，nums2[0] = 2 出现在 nums1 中下标 1 处，然而 nums2[1] = 4 和 nums2[2] = 6 没有出现在 nums2 中。因此，answer[1] = [4,6]。</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [1,2,3,3], nums2 = [1,1,2,2]
<strong>输出：</strong>[[3],[]]
<strong>解释：
</strong>对于 nums1 ，nums1[2] 和 nums1[3] 没有出现在 nums2 中。由于 nums1[2] == nums1[3] ，二者的值只需要在 answer[0] 中出现一次，故 answer[0] = [3]。
nums2 中的每个整数都在 nums1 中出现，因此，answer[1] = [] 。 
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums1.length, nums2.length &lt;= 1000</code></li>
	<li><code>-1000 &lt;= nums1[i], nums2[i] &lt;= 1000</code></li>
</ul>

## 标签

 - Array 
 - Hash Table 

## 代码

```golang
func findDifference(nums1 []int, nums2 []int) [][]int {
    ans:=[][]int{
        []int{},
        []int{},
    }
    hash1 := make(map[int]struct{})    
    hash2 := make(map[int]struct{})    
    for _,n:=range nums1{
        hash1[n] = struct{}{}
    }
    for _,n:=range nums2{
        if _,ok:=hash1[n];!ok{
            ans[1] = append(ans[1],n)
            // 去重
            hash1[n] = struct{}{}

        }
         hash2[n] = struct{}{}
    }
    for _,n:=range nums1{
        if _,ok:=hash2[n];!ok{
            ans[0] = append(ans[0],n)
            // 去重
            hash2[n] = struct{}{}
        }
    }
    return ans

}
```