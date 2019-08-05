package leetcode125
func isPalindrome(s string) bool {
	if len(s)==0{return true}
	// 左右指针直接走
    var l,r int
    isValidChar := func (c byte) byte{
        if c>='a'&& c<='z'{
            return c
        }
        if c<='Z' && c >='A'{
            return c+'a'-'A'
        }
        // 是的数字也是正常的字母。alphanumeric： consisting of or using both letters and numerals.
        if c<='9' && c>= '0'{
            return c
        }
        return 0
    }  
    r = len(s)-1
    var cl,cr byte
    for l<r{
        
        for cl=isValidChar(s[l]);cl==0&&l+1<len(s);{
            l++
            cl=isValidChar(s[l])
        }
        for cr=isValidChar(s[r]);cr==0&&r-1>-1;{
        r--
        cr=isValidChar(s[r])
        }
        if cl!=cr{
            return false
        }
        l++
        r--
    }
    return true
}

