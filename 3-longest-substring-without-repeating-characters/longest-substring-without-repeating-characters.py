class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        longest = 0
        length = len(s)
        got = {}
        i = 0
        c = 's'
        while(i<length):
            c = s[i]
            if c not in got:
                got[c] = i
                i+=1
            else :
                longest = max(len(got),longest);
                i = got[c]+1
                got.clear()

        return max(len(got),longest)
