class Solution {
    public int chalkReplacer(int[] chalk, int k) {
        // 前缀和
        int[] chalkAcc = new int[chalk.length]; 
        chalkAcc[0]=chalk[0];
        
        for(int i =1;i<chalk.length;++i){
            if(chalkAcc[i-1]>k)return i-1;
            chalkAcc[i]=chalkAcc[i-1]+ chalk[i];
        }
        
        k %= chalkAcc[chalk.length-1];
        // 这里用二分查找会更快
        for(int i = 0;i<chalk.length;++i)
        {
            if(chalkAcc[i]>k)return i;
        }
        
       return 0;

    }
}
