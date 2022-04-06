# 剑指 Offer II 031.最近最少使用缓存

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/OrIXps/](https://leetcode-cn.com/problems/OrIXps/)

## 题目描述

<div class="title__3Vvk">
<p>运用所掌握的数据结构，设计和实现一个&nbsp; <a href="https://baike.baidu.com/item/LRU" target="_blank">LRU (Least Recently Used，最近最少使用) 缓存机制</a> 。</p>

<p>实现 <code>LRUCache</code> 类：</p>

<ul>
	<li><code>LRUCache(int capacity)</code> 以正整数作为容量&nbsp;<code>capacity</code> 初始化 LRU 缓存</li>
	<li><code>int get(int key)</code> 如果关键字 <code>key</code> 存在于缓存中，则返回关键字的值，否则返回 <code>-1</code> 。</li>
	<li><code>void put(int key, int value)</code>&nbsp;如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。</li>
</ul>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre>
<strong>输入</strong>
[&quot;LRUCache&quot;, &quot;put&quot;, &quot;put&quot;, &quot;get&quot;, &quot;put&quot;, &quot;get&quot;, &quot;put&quot;, &quot;get&quot;, &quot;get&quot;, &quot;get&quot;]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
<strong>输出</strong>
[null, null, null, 1, null, -1, null, -1, 3, 4]

<strong>解释</strong>
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= capacity &lt;= 3000</code></li>
	<li><code>0 &lt;= key &lt;= 10000</code></li>
	<li><code>0 &lt;= value &lt;= 10<sup>5</sup></code></li>
	<li>最多调用 <code>2 * 10<sup>5</sup></code> 次 <code>get</code> 和 <code>put</code></li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>进阶</strong>：是否可以在&nbsp;<code>O(1)</code> 时间复杂度内完成这两种操作？</p>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 146&nbsp;题相同：<a href="https://leetcode-cn.com/problems/lru-cache/">https://leetcode-cn.com/problems/lru-cache/</a>&nbsp;</p>

## 标签

 - Design 
 - Hash Table 
 - Linked List 
 - Doubly-Linked List 

## 代码

```golang
type Node struct{
    Key int
    Val int
    Next *Node
    Prev *Node
}

type LRUCache struct {
    capacity int
    length int
    data *Node
    keyMap map[int]*Node
}
// 1.  实现LRU首先要维护有序，这需要一个有序的序列，
//   其顺序是最近使用的在前面，最进没使用的在后面 --- 可以考虑带头尾指针的链表 或者用 数组
//  插入的时候插入头部，如果满了就把最后一个删掉在插入，其他节点按顺序往后移
//  查询的时候将该节点交换到头部。其他按顺序往后移
//  考虑到这里发现用循环双链表会比较方便的实现插入和删除的同时实现有序，不需要移动其他节点
// 2. 要求O(1) 的查询和插入
// 仅仅是单链表肯定是没办法实现的，查询需要O(n) ，因此可以考虑加一个哈希表索引去掉查询的复杂度
// Constructor 
func Constructor(capacity int) LRUCache {
    
    return LRUCache{capacity:capacity,keyMap:make(map[int]*Node)}
}


func (this *LRUCache) Get(key int) int {
    ptr:= this.keyMap[key]
    if ptr == nil{
        return -1
    }
    if this.data!=ptr{
        // 1. ptr节点断开
        Detach(ptr)
        // 2. 将ptr插入头节点
        this.data = InsertHead(this.data,ptr)
    }
    
    return ptr.Val
}


func (this *LRUCache) Put(key int, value int)  {
    // 尝试插入，查找是否存在
    ptr:= this.keyMap[key]
    if ptr != nil{
        // 原来就有,直接修改内容
        ptr.Val= value
        // 除了头结点以外的交换到最前
        if this.data!=ptr{
            // 1. ptr节点断开
            Detach(ptr)
            // 2. 将ptr插入头节点
            this.data = InsertHead(this.data,ptr)
            // this.length
        }
        

    }else{
        // 不存在，添加或找一个最后的替换掉
        
        if this.length < this.capacity {
            // 1. 添加直接加在头节点
            node:= &Node{Key:key,Val:value}
            this.data = InsertHead(this.data,node)
            this.keyMap[key] = node
            this.length++
           
        }else{
            // fmt.Println(key,value,this,this.data)
            // 2. 找最后一个替换, 再将指针前移
            last:= this.data.Prev
            // fmt.Println("last",this.data.Prev)
            this.keyMap[last.Key] = nil

            last.Key = key
            last.Val = value
            this.data = last
            this.keyMap[key] = last
        } 
    }

}
// Detach 断开节点
func Detach(node  *Node) (*Node){
    // node 不应该有next和prev,应该是一个单纯的值节点
    if node == nil{
        // 空节点或者单个节点无法deatch
        return node
    }

    last:= node.Prev
    next:= node.Next
    last.Next = next
    next.Prev = last
    return node
}

// InsertHead 头插节点
func InsertHead(head  *Node,node  *Node) (*Node){

    // node 不应该有next和prev,应该是一个单纯的值节点
    if head == nil{
        node.Next = node
        node.Prev = node
        return node
    }

    prev := head.Prev
    head.Prev = node
    prev.Next = node
    node.Prev = prev
    node.Next = head
    return node
}

func (node *Node) String() (str string){
    str = ""
    head:=node
    for node!=nil&&node.Next!=head{
        str+=fmt.Sprintf("{%v,%v},",node.Key,node.Val)
        node = node.Next
    }
    return str
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
```