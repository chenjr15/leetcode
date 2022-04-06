# 剑指 Offer II 042.最近请求次数

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/H8086Q/](https://leetcode-cn.com/problems/H8086Q/)

## 题目描述

<p>写一个&nbsp;<code>RecentCounter</code>&nbsp;类来计算特定时间范围内最近的请求。</p>

<p>请实现 <code>RecentCounter</code> 类：</p>

<ul>
	<li><code>RecentCounter()</code> 初始化计数器，请求数为 0 。</li>
	<li><code>int ping(int t)</code> 在时间 <code>t</code> 添加一个新请求，其中 <code>t</code> 表示以毫秒为单位的某个时间，并返回过去 <code>3000</code> 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 <code>[t-3000, t]</code> 内发生的请求数。</li>
</ul>

<p><strong>保证</strong> 每次对 <code>ping</code> 的调用都使用比之前更大的 <code>t</code> 值。</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre>
<strong>输入：</strong>
inputs = [&quot;RecentCounter&quot;, &quot;ping&quot;, &quot;ping&quot;, &quot;ping&quot;, &quot;ping&quot;]
inputs = [[], [1], [100], [3001], [3002]]
<strong>输出：</strong>
[null, 1, 2, 3, 3]

<strong>解释：</strong>
RecentCounter recentCounter = new RecentCounter();
recentCounter.ping(1);     // requests = [<strong>1</strong>]，范围是 [-2999,1]，返回 1
recentCounter.ping(100);   // requests = [<strong>1</strong>, <strong>100</strong>]，范围是 [-2900,100]，返回 2
recentCounter.ping(3001);  // requests = [<strong>1</strong>, <strong>100</strong>, <strong>3001</strong>]，范围是 [1,3001]，返回 3
recentCounter.ping(3002);  // requests = [1, <strong>100</strong>, <strong>3001</strong>, <strong>3002</strong>]，范围是 [2,3002]，返回 3
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= t &lt;= 10<sup>9</sup></code></li>
	<li>保证每次对 <code>ping</code> 调用所使用的 <code>t</code> 值都 <strong>严格递增</strong></li>
	<li>至多调用 <code>ping</code> 方法 <code>10<sup>4</sup></code> 次</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 933&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/number-of-recent-calls/">https://leetcode-cn.com/problems/number-of-recent-calls/</a></p>

## 标签

 - Design 
 - Queue 
 - Data Stream 

## 代码

```golang
type RecentCounter struct {
    reqs []int
    // 指向下一个
    head int
    // 指向最后一个
    tail int
    cnt int  

}

func (this *RecentCounter) Tail() int{
    return this.reqs[this.tail]
}

func (this *RecentCounter) Dequeue() (val int){
    if this.cnt >0 {
        val = this.reqs[this.tail]
        this.tail++
        this.tail%=3000
        this.cnt--
    }
    return val
}
func (this *RecentCounter) Enqueue(val int) {
    if this.cnt>=len(this.reqs){
       this.Dequeue()
    }
    this.reqs[this.head]=val
    this.head++
    this.head %= 3000
    this.cnt++

}

func Constructor() RecentCounter {
    // 记录过去3000毫秒内发生的请求，每次的调用给的时间都是严格递增，
    // 那么可以用一个数组来记录这些请求，然后每次调用的时候就回去找数组里面所有时间大于等于当前t-3000的请求。
    // 因为这个数组是递增的，那么上面方法看起来像一个二分查找的应用。
    // 但是这里考虑到指定时间内没有重复请求，因此可以考虑用一个长度为3000的循环队列来存储请求
    // 这样数组的长度就是定长的了，然后每次t来了的时候就一个个出队直到队头的元素大于等于t-3000.
    // 更近一步可以直接用二分查找来找这个位置，然后直接移动对头指针。
    return RecentCounter{reqs:make([]int,3000)}
}


func (this *RecentCounter) Ping(t int) (cnt int) {
    lastime:= t-3000
    // 所有小于lasttime的都出队，
    for this.cnt>0 && this.Tail()  <lastime {
        this.Dequeue()
    }

  
    cnt = 1+this.cnt
    
    this.Enqueue(t)
    return cnt
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
```