class Solution {
    public int findMaximizedCapital(int k, int w, int[] profits, int[] capital) {
        // 贪心，每次在可以投资的项目中选择利润最高的
        // 1. 如何基于capital和profits排序
        // 最多投资的数目一定是小于项目数和限制数的
        int maxTimes = profits.length>k?k:profits.length;
        for(int times =0;times<maxTimes;++times){
            // 寻找可以投资的项目
            int maxI = 0;
            for(int i =0 ;i<profits.length){
                if(capital[i] <0){
                    continue;
                }
                maxI = i;
                if(capital[i]>w){
                    break;
                }
            }
            w += profits[maxI];
            // 设置标志位
            profits[maxI] = -1;
        }
        return w;
        

    }
}
