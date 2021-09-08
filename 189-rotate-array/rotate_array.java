class Solution {

    public void rotate(int[] nums, int k) {

        k=k%nums.length;
        if(k==0){
            return ;
        }
        int len= nums.length;
        // i -> (i+k)%len = j
        // j -> (j+len-k)%len 
        int t = nums[0];
        int j = len/2;
        int i = 0;
        int times = 0;
        // if((len%k) == 0) times = k;
        for(int start = 0;times<len;++start){
            i=start;
            t = nums[start];
            j = -1;
            while (j!=start){
                j = (i+len-k) %len;
                nums[i] = nums[j];
                i = j;
                times++;
            }
            nums[(start+k)%len] = t;
        }
        return ;
    }
}
