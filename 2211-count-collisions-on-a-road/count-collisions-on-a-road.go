func countCollisions(directions string) int {
    // stack := []int{}
    cnt :=0
    haveStay := 0
    haveGoRight := 0
    for _,dir:=range directions{

        if dir == 'L'{
            // 向左，判断一下栈是不是有，有就会撞
            if haveGoRight>0 {
                // 碰撞一次后停止第一次是+2，左边向右的也开始碰撞, cnt= 2 + haveGoRight-1 
                cnt+=haveGoRight+1
                haveGoRight = 0
                haveStay = 1
            }else if haveStay>0 {
                cnt++
            }
        }else if dir == 'R'{
            
            haveGoRight++
            
        }else if dir == 'S'{
            if haveGoRight>0 {
                // 碰撞一次后停止，左边向右的也开始碰撞, cnt= 1 + haveGoRight-1 
                cnt+=haveGoRight
                haveGoRight = 0
            }
            haveStay = 1
        }
    }
    return cnt

}