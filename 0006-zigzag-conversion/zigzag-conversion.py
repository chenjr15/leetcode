class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows ==1 :
            return s
        results = ["" for i in range(numRows)]
        block_size = numRows-2 +numRows
        for i in range(len(s)):
            in_block =i%block_size
            loc_row = in_block
            if in_block >= numRows:
                loc_row = numRows - (in_block - numRows +1)-1
            results[loc_row]+=s[i]
        return "".join(results)