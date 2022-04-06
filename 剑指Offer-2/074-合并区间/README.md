# 剑指 Offer II 074.合并区间

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/SsGoHC/](https://leetcode-cn.com/problems/SsGoHC/)

## 题目描述

<p>以数组 <code>intervals</code> 表示若干个区间的集合，其中单个区间为 <code>intervals[i] = [start<sub>i</sub>, end<sub>i</sub>]</code> 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>intervals = [[1,3],[2,6],[8,10],[15,18]]
<strong>输出：</strong>[[1,6],[8,10],[15,18]]
<strong>解释：</strong>区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
</pre>

<p><strong>示例&nbsp;2：</strong></p>

<pre>
<strong>输入：</strong>intervals = [[1,4],[4,5]]
<strong>输出：</strong>[[1,5]]
<strong>解释：</strong>区间 [1,4] 和 [4,5] 可被视为重叠区间。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= intervals.length &lt;= 10<sup>4</sup></code></li>
	<li><code>intervals[i].length == 2</code></li>
	<li><code>0 &lt;= start<sub>i</sub> &lt;= end<sub>i</sub> &lt;= 10<sup>4</sup></code></li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 56&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/merge-intervals/">https://leetcode-cn.com/problems/merge-intervals/</a></p>

## 标签

 - Array 
 - Sorting 

## 代码

```java
class Solution {
    public int[][] merge(int[][] intervals) {
        ArrayList<int[]> merged = new ArrayList<>();
        
        Arrays.sort(intervals,(a,b)->a[0]==b[0]?b[1]-a[1]:a[0]-b[0]);
        // System.out.println(intervals);
        int start,end;
        start = intervals[0][0];
        end = intervals[0][1];
        for(int i =1;i<intervals.length;++i){
            if (intervals[i][0]<=end){
                // 起点在现有区间的终点结束之前，因此肯定要合并
                // 判断其结尾是否扩大当前区间。
                if (intervals[i][1]>end){
                    end = intervals[i][1];
                }
            }else{
                //  新区间
                merged.add(new int[]{start,end});
                start = intervals[i][0];
                end = intervals[i][1];
            }

        }
        merged.add(new int[]{start,end});
        
        return merged.toArray(new int[][]{} );

    }
}
```