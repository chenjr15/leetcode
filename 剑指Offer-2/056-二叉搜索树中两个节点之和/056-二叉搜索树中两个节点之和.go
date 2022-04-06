/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    // 尝试模仿双指针
    var sleft,sright []*TreeNode
    // fmt.Println(sright)
    pleft := root
    pright := root
    for pleft!=nil{
        sleft= append(sleft,pleft)
        pleft = pleft.Left
    }
    goRight := func ()(r *TreeNode){
        if len(sleft)==0{
            return nil
        }
        pleft = sleft[len(sleft)-1]
        sleft = sleft[:len(sleft)-1]
        r = pleft
        pleft = pleft.Right
        for pleft!=nil{
            sleft= append(sleft,pleft)
            pleft = pleft.Left
        }
        return r
    }
    for pright!=nil{
        sright= append(sright,pright)
        pright = pright.Right
    }
    // 从右往左走
    goLeft := func ()(r *TreeNode){
        if len(sright)==0{
            return nil
        }
        pright = sright[len(sright)-1]
        sright = sright[:len(sright)-1]
        r = pright
        pright = pright.Left
        for pright!=nil{
            sright= append(sright,pright)
            pright = pright.Right
        }
        return r
    }

    // 左中右，中序遍历的非递归实现
    // for {
    //     p:= goRight()
    //     if p==nil{
    //         break
    //     }
    //     fmt.Println(p.Val)
    // }
    // for {
    //     p:= goLeft()
    //     if p==nil{
    //         break
    //     }
    //     fmt.Println(p.Val)
    // }
    var l,r *TreeNode
    
    l=goRight()
    r=goLeft()
    for l!=r{
        sum:=l.Val+r.Val
        if sum == k{
            return true
        }else if sum>k{
            r = goLeft()
        }else{
            l = goRight()
        }
    }
  

    return false

}

