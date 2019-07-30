package leetcode28
// strStr Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
func strStr(haystack string, needle string) int {
    if needle == ""{
        return 0
    }
    if len(needle)> len(haystack){
        return -1
    }
    
    for i:=0;i<=len([]byte(haystack)) - len([]byte(needle));i++ {
        p :=0 
        for ; p<len(needle) && needle[p] == haystack[i+p];p++{
        }
        if p == len(needle){
           return i 
        }
    }
    return -1
}
