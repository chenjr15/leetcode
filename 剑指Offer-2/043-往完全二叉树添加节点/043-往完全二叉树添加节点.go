/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
    // 找到倒数第二层，找倒数第二层最后一个
    // 需要保存  最后一个的节点
    levelCnt      []int 
    nodeQueue     []*TreeNode
    curLevelCnt  int 
    root *TreeNode

}


func Constructor(root *TreeNode) CBTInserter {
    inst := CBTInserter{
        levelCnt:     []int{},
        nodeQueue:     []*TreeNode{root},
        curLevelCnt :0,
        root :root,
    }
    
    if root==nil{
        // 输入情况以外
        return inst
    }
    // 遍历树 
    queue := inst.nodeQueue[:]
    start:=0
    curLevel:=0
    nextLevelCnt:=0
    curLevelCnt:=1
    // 层次遍历
    for ;start<len(queue);start++{
        if curLevelCnt == 0 {
            curLevelCnt = nextLevelCnt
            nextLevelCnt=0
            curLevel++
        }
        // 取出头
        node:= queue[start]
        // fmt.Printf("level %d : %v\n",curLevel,node.Val)
        curLevelCnt--
        
        // queue = queue[:len(queue)-1]
        // 加入队列
        // inst.nodeQueue = append(inst.nodeQueue,root)
        // 左右节点加入队列
        if node.Left!=nil{
            queue = append(queue,node.Left)
            nextLevelCnt++
        }
        if node.Right!=nil{
            queue = append(queue,node.Right)
            nextLevelCnt++
        }

    }
    inst.nodeQueue = queue
    // fmt.Println(inst.nodeQueue,queue)
    return inst
}


func (this *CBTInserter) Insert(v int) int {
    node := &TreeNode{Val:v}
    this.nodeQueue = append(this.nodeQueue,node)
    father:= this.nodeQueue[len(this.nodeQueue)/2-1]
    // fmt.Println("father of ",v,len(this.nodeQueue)/2-1,father,this.nodeQueue)

    if father.Left!=nil{
        father.Right = node
    }else {
        father.Left = node
    }
    return father.Val
}


func (this *CBTInserter) Get_root() *TreeNode {
    return this.root
}


/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */