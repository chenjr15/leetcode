func dailyTemperatures(temperatures []int) []int {
    // 单调栈
    // 其实就是用一个栈里面放着所有小于当前温度的东西，如何保证所有都小于呢，那就是碰到比栈顶大的就放栈顶出栈
    // 这里会保证栈单调递减
    result:=make([]int,len(temperatures))
    stack:=make([]int,0,len(temperatures))
    stacktop:=0
    for i,t:= range temperatures{

        // 出栈并计算索引差值
        for  len(stack)>0 {
            stacktop =stack[len(stack)-1];
            if  temperatures[stacktop]< t{
                result[stacktop] = i-stacktop
                // 出栈
                stack=stack[:len(stack)-1]
            }else{
                break
            }
        }
        
        stack=append(stack,i)
        
    }
    return result

}