func optimalDivision(nums []int) string {

    builder:=strings.Builder{}
    n:=len(nums)
    if n>0{
        i:=0
        // 头
        builder.WriteString(fmt.Sprintf("%d",nums[i]))
        i++
        if n>1{
           
            builder.WriteByte('/')
            if n>2{

                builder.WriteByte('(')
                
            }
            builder.WriteString(fmt.Sprintf("%d",nums[i]))
        }
        // 2 -- n-1
        for i=2;i<n;i++{
            builder.WriteByte('/')
            builder.WriteString(fmt.Sprintf("%d",nums[i]))
        }
         if n>2{
            builder.WriteByte(')')
        }
    }
        
       

    
    return builder.String()

}