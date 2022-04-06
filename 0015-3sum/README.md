# 15.三数之和

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/3sum/](https://leetcode-cn.com/problems/3sum/)

## 题目描述

<p>给你一个包含 <code>n</code> 个整数的数组 <code>nums</code>，判断 <code>nums</code> 中是否存在三个元素 <em>a，b，c ，</em>使得 <em>a + b + c = </em>0 ？请你找出所有和为 <code>0</code> 且不重复的三元组。</p>

<p><strong>注意：</strong>答案中不可以包含重复的三元组。</p>

<p> </p>

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

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 <= nums.length <= 3000</code></li>
	<li><code>-10<sup>5</sup> <= nums[i] <= 10<sup>5</sup></code></li>
</ul>

## 标签

 - Array 
 - Two Pointers 
 - Sorting 

## 代码

```golang
func threeSum(nums []int) [][]int {
    lenNums := len(nums)
    results:=make([][]int,0,100)
    if lenNums<3{
        return results
    }
    // 固定一个元素，搜索剩下两个，变成一个twosum问题
    // 将元素排序，加快后面的twoSum搜索
    sort.Ints(nums)
    var pb,pc,b,c,tempSum int
    
    for pa,a := range nums{
        if a>0{
            // 那么后面一定也大于0了，所有大于0的数字加起来不会等于0 的
            break
        }
        if pa>0 && nums[pa] == nums[pa-1]{
            continue
        }
        // a+b+c = 0 ==> b+c = -a
        pb,pc=pa+1,lenNums-1
        // 双指针搜索,注意这里会导致每一个都会搜索一次
        // 这里要注意因为有重复元素，所以要注意去重
        a=-a
        for pb<pc {
            b = nums[pb]
            c = nums[pc]
            tempSum = b+c
            
            if tempSum == a{
                results = append(results,[]int{-a,b,c})
                pc--
                pb++
                for pb<pc&& nums[pb] == b {
                    pb++
                }  
                
                for pb<pc && nums[pc] == c {
                    pc-- 
                }
            }else if tempSum< a {
                // for pb<pc&& nums[pb] == b {
                    pb++
                // }  
            }else{
                // for pb<pc && nums[pc] == c {
                    pc-- 
                // }
            }
        }
    }
    return results
}
```