/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    stack []*TreeNode
    
}
func stack2str (nodes []*TreeNode)string{
    if len(nodes) ==0 {
        return ""
    }
    builder := strings.Builder{}
    for i:=0;i<len(nodes);i++{
        builder.WriteString(strconv.Itoa(i))
        builder.WriteRune(':')
        builder.WriteString(strconv.Itoa(nodes[i].Val))
        builder.WriteRune(',')
    }
    // builder.WriteString(strconv.Itoa(nodes[len(nodes)-1].Val))
    return builder.String()
}

func Constructor(root *TreeNode) BSTIterator {
    stack:= []*TreeNode{}
    p := root
    // --- 标准非递归中序遍历 ---
    // for p !=nil || len(stack)!=0{
    //     // 一直往左遍历到当前的最左节点，并全部入栈
    //     for p!=nil{
    //         stack = append(stack,p)
    //         p = p.Left
    //     }
    //     if len(stack)>0 {
    //         // 出栈
    //         p = stack[len(stack)-1]
    //         stack = stack[:len(stack)-1]
    //         // 出栈的时候就是要访问的时候
    //         // visit(p)
    //         fmt.Println(p)
    //         p = p.Right
    //     }
    // }
    // 一直往左遍历到当前的最左节点，并全部入栈
    for p!=nil{
            stack = append(stack,p)
            p = p.Left
    }
    return BSTIterator{stack}

    
}


func (this *BSTIterator) Next() (nextval int) {
    // 题目说保证有效，暂时不管stack是否为空
    // fmt.Println("before next",stack2str(this.stack))
    p := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    nextval = p.Val
    p = p.Right
    for p!=nil{
        this.stack = append(this.stack,p)
        p = p.Left
    }
    // fmt.Println("after ",nextval,stack2str(this.stack))
    return nextval
}   


func (this *BSTIterator) HasNext() bool {
    return len(this.stack)!=0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */