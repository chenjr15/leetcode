# 剑指 Offer II 007.数组中和为 0 的三个数

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/1fGaJU/](https://leetcode-cn.com/problems/1fGaJU/)

## 题目描述

<p>给定一个包含 <code>n</code> 个整数的数组&nbsp;<code>nums</code>，判断&nbsp;<code>nums</code>&nbsp;中是否存在三个元素&nbsp;<code>a</code> ，<code>b</code> ，<code>c</code> <em>，</em>使得&nbsp;<code>a + b + c = 0</code> ？请找出所有和为 <code>0</code> 且&nbsp;<strong>不重复&nbsp;</strong>的三元组。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [-1,0,1,2,-1,-4]
<strong>输出：</strong>[[-1,-1,2],[-1,0,1]]
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = []
<strong>输出：</strong>[]
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums = [0]
<strong>输出：</strong>[]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= nums.length &lt;= 3000</code></li>
	<li><code>-10<sup>5</sup> &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 15&nbsp;题相同：<a href="https://leetcode-cn.com/problems/3sum/">https://leetcode-cn.com/problems/3sum/</a></p>

## 标签

 - Array 
 - Two Pointers 
 - Sorting 

## 代码

```golang
func threeSum(nums []int) [][]int {
    triples := make([][]int,0)
    length := len(nums)
    if length <3 {
        return triples
    }
    // 排序
    sort.Ints(nums)
    var idxB,idxC,target int
    for idxA,a :=range nums{
        // 如果a 大于0 则直接break，因为后面的所有都是正数再加也不会等于0了
        if a>0{
            break
        }
        // 对第一个数去重
        if idxA > 0 && nums[idxA]== nums[idxA-1]{
            continue
        }
        target = -a
        // 双指针搜索TwoSum 同时去重
        idxB=idxA+1
        idxC=length-1
        for idxB<idxC{
            
            if (nums[idxB] + nums[idxC]) == target{
                triples = append(triples,[]int{a,nums[idxB],nums[idxC]})
                // 数组有序且要求答案不重复，因此两个指针一定要前后移动
                idxB++
                idxC--
                for idxB<idxC && nums[idxB] == nums[idxB-1]{
                    idxB++
                }
                for idxB<idxC && nums[idxC] == nums[idxC+1]{
                    idxC--
                }
                
            }else if (nums[idxB] + nums[idxC]) > target{
                idxC--
            }else{
                idxB++
            }
        }
        

    }
    return triples

}
```