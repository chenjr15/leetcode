func convertToTitle(columnNumber int) string {
    r:=0
    runes:=[]rune{}
    
    // columnNumber--
    for columnNumber>0{
        columnNumber--
        r =  columnNumber%26
        columnNumber/=26
        runes  = append(runes,rune(r+'A'))
    }
    // runes  = append(runes,rune(columnNumber+'A'))
    // fmt.Println(runes)
    // reverse
    for i,j:=0,len(runes)-1;i<j;i,j=i+1,j-1{
        runes[i],runes[j] = runes[j],runes[i]
    }
    return string(runes)
}