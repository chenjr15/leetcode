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