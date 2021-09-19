class Solution {
    public int minSteps(int n) {
        // 处理可以直接出结果的n
        if(n ==1)return 0;
        if(n<6)return n;
        int step=0;
        
        for(int i =2;i<=n/2&&n>5;++i){
            if((n%i)!=0){
                // 不能被i整除
                continue;
            }
            // System.out.println(""+n+"/"+i+"="+n/i);
            // 被i整除
            step+= i;
            n=n/i;
            i=1;
        }
        // System.out.println(n);
        // 得到质数
        step+=n;
        return step;
    }
}
