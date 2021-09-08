class Solution {
    public int balancedStringSplit(String s) {
        // 有一些像括号匹配
        int count = 0;
        // 记录l的个数，如果碰到字符是L就是lCount +1 ,否则就是lCount-1
        // 这样如果平衡了lCount就是0，因此判断lCount是0就说明已经平衡了。
        // 题目的 最大数量就是每个平衡字符串最小，因此碰到了直接计数+1 即可
        int lCount  =  0;
        for(int i =0;i<s.length();++i){
            lCount +=(s.charAt(i) == 'L')?1:-1;
            if(lCount==0){
                ++count;
            }
        }
        return count;

    }
}
