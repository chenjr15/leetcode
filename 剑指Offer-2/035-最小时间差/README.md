# 剑指 Offer II 035.最小时间差

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/569nqc/](https://leetcode-cn.com/problems/569nqc/)

## 题目描述

<p>给定一个 24 小时制（小时:分钟 <strong>&quot;HH:MM&quot;</strong>）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>timePoints = [&quot;23:59&quot;,&quot;00:00&quot;]
<strong>输出：</strong>1
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>timePoints = [&quot;00:00&quot;,&quot;23:59&quot;,&quot;00:00&quot;]
<strong>输出：</strong>0
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= timePoints &lt;= 2 * 10<sup>4</sup></code></li>
	<li><code>timePoints[i]</code> 格式为 <strong>&quot;HH:MM&quot;</strong></li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 539&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/minimum-time-difference/">https://leetcode-cn.com/problems/minimum-time-difference/</a></p>

## 标签

 - Array 
 - Math 
 - String 
 - Sorting 

## 代码

```golang
func findMinDifference(timePoints []string) int {
    minDiff := 1440
    if len(timePoints)<2{
        return minDiff
    }
    minutes := make([]int,len(timePoints))
    for i,ts := range timePoints{
        // 转换为分钟数
        m := (int(ts[0]-'0')*10+int(ts[1]-'0'))*60+int(ts[3]-'0')*10+int(ts[4]-'0')
        // fmt.Println(ts,m)
        minutes[i] = m
    }
    sort.Ints(minutes)
    minDiff =  minDiff - minutes[len(minutes)-1] +  minutes[0]
    for i:=0;i<len(minutes)-1;i++{
        if minDiff <0|| minDiff> minutes[i+1] -minutes[i]{
            minDiff =  minutes[i+1] -minutes[i]
        }
        
    }
    return minDiff
}
```