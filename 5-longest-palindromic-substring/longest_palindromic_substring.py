class Solution:
        
        
    def outter_exam(self,s:str):
        if (len(s)==1):
            return True
        elif len(s)>1:
            return s[0] == s[-1]
    def expend_exam(self,s:str,center:int):
        left = center
        right = center +1
        while  right<(len(s)):
            # exam expend to both side first
            if left > 0 and self.outter_exam(s[left-1:right+1]):
                right+=1
                left -=1
                continue

            # fail than expend to right only
            elif s[center] == s[left] and self.outter_exam(s[left:right+1]):
                right +=1
                continue
            else:
               break
        #print("reutring %d:%d %s "%(left,right,s[left:right]))

        return s[left:right]
    
    def longestPalindrome(self, s: str) -> str:
        longest_str = ""
        for i in range(len(s)):
    
            expended = self.expend_exam(s,i)
            if len(expended)>len(longest_str):
                longest_str = expended
    
        return longest_str
