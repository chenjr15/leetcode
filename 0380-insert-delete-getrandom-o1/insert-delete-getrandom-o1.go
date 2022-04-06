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