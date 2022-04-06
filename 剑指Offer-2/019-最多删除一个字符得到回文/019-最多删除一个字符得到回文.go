func validPalindrome(s string) bool {
    // 题目说只改一个字符，因此我们在碰到不是回文的字符的情况下，
    // 尝试将左边或者右边删掉，看是否能够形成回文川，也就是只要试两次
    left:=0
    right:=len(s)-1
    for left<right{

        // fmt.Printf("%d,%d,%c,%c\n",left,right,lchar,rchar)
        if s[left] != s[right]{
            // 尝试判断左边或者右边删掉的情况
            // 删掉左边
            if !isPalindrome(s,left+1,right){
                return isPalindrome(s,left,right-1)
            }else {
                return true
            }
        }
        right--
        left++
    }
    return true

}

func isPalindrome(s string, left ,right int ) bool{
    for left<right{
        // fmt.Printf("%d,%d,%c,%c\n",left,right,lchar,rchar)
        if s[left] != s[right]{
            // 尝试判断左边或者右边删掉的情况
            // 删掉左边
            return false 
        }
        right--
        left++
    }
    return true
}