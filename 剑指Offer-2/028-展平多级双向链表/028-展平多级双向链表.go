/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
    p:=root
    next:=root
    for p!=nil{
        if p.Child!=nil{
            // 存在子节点，找到子节点的横向尾巴一个(即不管孙节点)
            pChildTail := p.Child
            for pChildTail.Next!=nil{
                pChildTail=pChildTail.Next
            }
            // 将下级横向整个拼接到当前节点后面
            next = p.Next
            // 更新next指针
            p.Next = p.Child
            p.Child.Prev = p
            // 断开child
            p.Child = nil
            pChildTail.Next = next
            if next!=nil {
                // 当前节点没有下一个节点的情况
                next.Prev = pChildTail
            }
            
        }
        p=p.Next
    }

    return root
    
}