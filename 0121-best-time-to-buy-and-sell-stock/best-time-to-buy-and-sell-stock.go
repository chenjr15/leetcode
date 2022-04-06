func maxProfit(prices []int) int {
    if len(prices) ==0{
        return -1
    }
    min :=prices[0]

    diff :=0
    maxDiff :=0
    for _,n := range prices{
        diff = n-min
        if n<min{
            min = n
        }
        if diff >0 && diff > maxDiff{
            maxDiff = diff
        }
    }
    return maxDiff

}