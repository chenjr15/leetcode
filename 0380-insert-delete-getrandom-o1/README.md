# 380.O(1) 时间插入、删除和获取随机元素

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/insert-delete-getrandom-o1/](https://leetcode-cn.com/problems/insert-delete-getrandom-o1/)

## 题目描述

<p>实现<code>RandomizedSet</code> 类：</p>

<div class="original__bRMd">
<div>
<ul>
	<li><code>RandomizedSet()</code> 初始化 <code>RandomizedSet</code> 对象</li>
	<li><code>bool insert(int val)</code> 当元素 <code>val</code> 不存在时，向集合中插入该项，并返回 <code>true</code> ；否则，返回 <code>false</code> 。</li>
	<li><code>bool remove(int val)</code> 当元素 <code>val</code> 存在时，从集合中移除该项，并返回 <code>true</code> ；否则，返回 <code>false</code> 。</li>
	<li><code>int getRandom()</code> 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 <strong>相同的概率</strong> 被返回。</li>
</ul>

<p>你必须实现类的所有函数，并满足每个函数的 <strong>平均</strong> 时间复杂度为 <code>O(1)</code> 。</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre>
<strong>输入</strong>
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
<strong>输出</strong>
[null, true, false, true, 2, true, false, 2]

<strong>解释</strong>
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>-2<sup>31</sup> &lt;= val &lt;= 2<sup>31</sup> - 1</code></li>
	<li>最多调用 <code>insert</code>、<code>remove</code> 和 <code>getRandom</code> 函数 <code>2 *&nbsp;</code><code>10<sup>5</sup></code> 次</li>
	<li>在调用 <code>getRandom</code> 方法时，数据结构中 <strong>至少存在一个</strong> 元素。</li>
</ul>
</div>
</div>

## 标签

 - Design 
 - Array 
 - Hash Table 
 - Math 
 - Randomized 

## 代码

```golang
type RandomizedSet struct {
    data []int
    idxHash map[int]int
    

}

// 用一个哈希表来存储数据，可以做到O(1)的查找和插入删除，但是没办法做到O(1)的随机读取
// 考虑加上一个data数组，data数组保存所有的数据，但是他是无序的，用来做随机读取
// 用哈希表保存每个数字在data中的下标，每次插入就插入在尾部，实现O(1)的插入
// 在删除的时候则是用最后面地元素交换到要删的元素的位置，再讲数组的长度缩减1 实现O(1)的删除

func Constructor() RandomizedSet {
    return RandomizedSet{data:make([]int,0),idxHash:make(map[int]int)}
}


func (this *RandomizedSet) Insert(val int) bool {
    if pos := this.idxHash[val];pos !=0{
        return false
    }
    this.data = append(this.data,val)
    this.idxHash[val]= len(this.data)

    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if len(this.data) == 0{
        return false
    }
    // 尝试查找下标
    pos := this.idxHash[val]
    // 下标不存在直接返回失败
    if pos==0{
        return false
    }

    
    // 用最后一个替换掉当前要删掉的位置的东西，
    lastone := this.data[len(this.data)-1]
    this.data[pos-1]= lastone
    // 更新原来最后一个的索引
    this.idxHash[lastone]=pos
    // 然后将原最后一个的位置删掉
    this.data = this.data[0:len(this.data)-1]
    // 清除原来的索引
    this.idxHash[val] = 0

    // fmt.Println("remove",val,pos ,this.data,this.idxHash)

    return true

}


func (this *RandomizedSet) GetRandom() int {
    if len(this.data) == 0{
        return 0
    }


    return this.data[rand.Intn(len(this.data))] 
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
```