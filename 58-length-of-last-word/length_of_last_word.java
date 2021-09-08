class Solution {
    public int lengthOfLastWord(String s) {
        int len = 0;
        int strlen = s.length();
        strlen--;
        char c;
        for(;strlen>=0;--strlen){
            c = s.charAt(strlen);
            if(c == ' '){
                if(len != 0){
                    return len;
                }else{
                    continue;
                }
            }else{
                ++len;
            }
        }
        return len;

    }
}
