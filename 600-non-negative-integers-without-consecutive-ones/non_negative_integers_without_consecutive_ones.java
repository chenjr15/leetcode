class Solution {
    public int countBitlen(int n ){
        int bitlen = 0;
        if(n<0)return 0;
        if(n==0){return 1;}
        while((n>>bitlen) !=0){
            ++bitlen;
        }
        return bitlen;
    }
    public int findIntegers(int n) {
        int total = 0;
        int bitlen = countBitlen(n);
        if(bitlen==1){
            return 2;
        }
        int[] numdp= new int[bitlen];
        numdp[0]= 1;
        numdp[1] =2;
        for(int i=2;i<bitlen;++i){
            numdp[i] = numdp[i-1]+numdp[i-2];
        }
        while(bitlen>1){
            total += numdp[bitlen-1];
            if((n& (1<<(bitlen-2))) != 0){
                // 次高位不为0直接加上剩下的所有位数即可

                total += numdp[bitlen-2];
                n=0;
                bitlen=0;
            }else{
                // 削掉最高位和次高位
                n = n &( ~(11<<(bitlen-2)));
                bitlen = countBitlen(n);
                if(bitlen==1){
                    total+=n+1;
                }
            }
        }
        return total;

    }
}

