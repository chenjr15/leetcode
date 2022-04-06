func evalRPN(tokens []string) int {
    stack := make([]int,0)
    var val1,val2 int
    for _,tk := range tokens{

        if tk == "+"{
            // 出栈
            stack,val1,val2 = pop(stack)
            stack = append(stack,val1+val2)

        }else if tk == "-"{
            stack,val1,val2 = pop(stack)
            stack = append(stack,val1-val2)
        }else if tk == "*"{
            stack,val1,val2 = pop(stack)
            stack = append(stack,val1*val2)
        }else if tk == "/"{
            stack,val1,val2 = pop(stack)
            stack = append(stack,val1/val2)
        }else {
            // 入栈
            val1,err := strconv.Atoi(tk)
            if err!=nil{
                fmt.Println("Error",err,tk)
                return -1
            }
            stack = append(stack,val1)

        }
    }
    return stack[0]
}
func pop(stack []int) (s []int,val1,val2 int){
    val1,val2  = stack[len(stack)-2], stack[len(stack)-1]
    return stack[:len(stack)-2],val1,val2
}