func isPalindrome(s string) bool {
    left:=0
    right:=len(s)-1
    var rchar,lchar byte 
    for left<right{
        lchar = s[left]
        // 跳过符号
        for left<right && !isLetterNum(lchar) {
            // fmt.Printf("Passing: %c at %d %v\n",rune(lchar),left,isLetterNum(lchar))
            left++
            lchar = s[left]
        }
        rchar = s[right]
        for left<right && !isLetterNum(rchar) {
            // fmt.Printf("Passing: %c at %d %v\n",rune(rchar),right,isLetterNum(rchar))
            right--
            rchar = s[right]
        }
        rchar = toLowerCase(rchar)
        lchar = toLowerCase(lchar)
        // fmt.Printf("%d,%d,%c,%c\n",left,right,lchar,rchar)
        if rchar != lchar{
            return false
        }
        right--
        left++

    }
    return true

}

func toLowerCase(c byte) byte{
    if  (c>='A' && c <='Z'){
        c += 'a'-'A'
    }
    return c

}

func isLetterNum(c byte)bool{
    if  (c>='A' && c <='Z'){
        return true
    }
    if (c>='a' &&  c <='z'){
        return true
    }
    if (c>='0' &&  c <='9'){
        return true
    }
    return false

}