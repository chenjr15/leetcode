class Solution {
    public boolean checkValidString(String s) {
        // 左括号的最小数量
        int low = 0;
        // 左括号的最大数量
        int high =0;
        for(int i =0;i<s.length();++i ){
            switch (s.charAt(i)){
                case '(':
                    // 最大数量和最小数量都加1
                    low++;
                    high++;
                    break;
                case '*':
                    // 星号可以是左括号也可以是右括号或者是空，
                    high++;
                    if(low>0)
                    low--;
                    
                    
                    break;
                case ')':
                    // 碰到右括号，同时减掉
                    if(low>0)
                    low--;
                    high--;
                    if(high<0){
                        return false;
                    }
            }
        }
        // 下限一定不能小于0
        return low<=0;
    }
}
