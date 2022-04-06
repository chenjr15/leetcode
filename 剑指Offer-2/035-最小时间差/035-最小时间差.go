func findMinDifference(timePoints []string) int {
    minDiff := 1440
    if len(timePoints)<2{
        return minDiff
    }
    minutes := make([]int,len(timePoints))
    for i,ts := range timePoints{
        // 转换为分钟数
        m := (int(ts[0]-'0')*10+int(ts[1]-'0'))*60+int(ts[3]-'0')*10+int(ts[4]-'0')
        // fmt.Println(ts,m)
        minutes[i] = m
    }
    sort.Ints(minutes)
    minDiff =  minDiff - minutes[len(minutes)-1] +  minutes[0]
    for i:=0;i<len(minutes)-1;i++{
        if minDiff <0|| minDiff> minutes[i+1] -minutes[i]{
            minDiff =  minutes[i+1] -minutes[i]
        }
        
    }
    return minDiff
}