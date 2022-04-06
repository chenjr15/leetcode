# 剑指 Offer II 060.出现频率最高的 k 个数字

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/g5c51o/](https://leetcode-cn.com/problems/g5c51o/)

## 题目描述

<p>给定一个整数数组 <code>nums</code> 和一个整数 <code>k</code>&nbsp;，请返回其中出现频率前 <code>k</code> 高的元素。可以按 <strong>任意顺序</strong> 返回答案。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入: </strong>nums = [1,1,1,2,2,3], k = 2
<strong>输出: </strong>[1,2]
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入: </strong>nums = [1], k = 1
<strong>输出: </strong>[1]</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>k</code> 的取值范围是 <code>[1, 数组中不相同的元素的个数]</code></li>
	<li>题目数据保证答案唯一，换句话说，数组中前 <code>k</code> 个高频元素的集合是唯一的</li>
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>所设计算法的时间复杂度 <strong>必须</strong> 优于 <code>O(n log n)</code> ，其中 <code>n</code><em>&nbsp;</em>是数组大小。</p>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 347&nbsp;题相同：<a href="https://leetcode-cn.com/problems/top-k-frequent-elements/">https://leetcode-cn.com/problems/top-k-frequent-elements/</a></p>

## 标签

 - Array 
 - Hash Table 
 - Divide and Conquer 
 - Bucket Sort 
 - Counting 
 - Quickselect 
 - Sorting 
 - Heap (Priority Queue) 

## 代码

```java
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        // hashmap + 最小堆或者快排一部分
        HashMap<Integer,Integer> map = new HashMap<>();
        for(int n:nums){
            Integer cnt = map.get(n);
            if (cnt==null){
                cnt=1;
            }else{
                cnt++;
            }
            map.put(n,cnt);
        }
        // 默认小顶堆，即每次弹出最小的
        PriorityQueue<Map.Entry<Integer,Integer>> pq = new PriorityQueue<>((e1,e2) -> e1.getValue()- e2.getValue());
        for (Map.Entry < Integer, Integer > entry: map.entrySet()) {
//            System.out.println(pq.toString());
            if(pq.size()<k){
                pq.offer(entry);
            }else{
                if (entry.getValue()>pq.peek().getValue()){
                    pq.poll();
                    pq.offer(entry);
                }
            }
        }
        int minlen = Math.min(k,pq.size());
        int[] result = new int[minlen];

        for (int i = 0; i < minlen; i++) {
            result[i] = pq.poll().getKey();
        }
        return result;

    }
}
```