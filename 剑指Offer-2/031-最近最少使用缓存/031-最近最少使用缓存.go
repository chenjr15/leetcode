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