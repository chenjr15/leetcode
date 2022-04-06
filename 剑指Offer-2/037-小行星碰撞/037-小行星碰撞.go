func asteroidCollision(asteroids []int) []int {
    // 1. 符号相同的不会相遇
    // 2. 相遇后绝对值大的留下，相等的同时消失
    // 考虑用双链表模拟，先构建双链表，然后从做到右找第一个异号的数字开始碰
    // 碰了以后将消失的删掉，然后留下的在和左边的比符号，如果左边为空则和右边比符号，重复上述步骤，直到没有异号
    // 注意最左边的负号将没得碰，因此直接跳掉
    // -- 这个过程可以直接用一个栈来简化
    stack:= make([]int,0,len(asteroids))
    for _,star := range asteroids{
        if len(stack) == 0 || stack[len(stack)-1] < 0||star >0{
            // 左边是空或者全是符号的,或者当前行星是往右的，直接进栈
            stack=append(stack,star)
        }else if star*stack[len(stack)-1] <0{
            // fmt.Println("开始撞",stack,star)
            // 就一直撞直到：1. 左边没了，2.右边的撞没 3.左边的是往左走的
            for len(stack) != 0 && star!=0 && stack[len(stack)-1]>0{
                // 异号，且star<0 ,开始撞
                if -star == stack[len(stack)-1] {
                    // 同归于尽
                    stack=stack[:len(stack)-1]
                    // 消除star
                    star = 0 
                }else if -star < stack[len(stack)-1]{
                    // 右边消失
                    star = 0 
                    break
                }else{
                    // 左边消失，需要将左边的提上来再比
                    stack=stack[:len(stack)-1]
                }
            }
            // fmt.Println("结果",stack,star)
            if star != 0{
                stack=append(stack,star)
            }

        }
    }
    return stack

}