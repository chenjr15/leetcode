class Solution {

    public int max(int a,int b){
        return b>a?b:a;
    }
    public int lengthOfLIS(int[] nums) {

        int max_len = 1;

        // dp数组的含义是以下标i元素结尾的最长递增子序列的长度为dp[i]
        int[] dp = new int[nums.length];
        // 那么我们知道dp[0]肯定是1，因为只有一个元素
        // 那么dp[1] 就是看nums[1]是不是大于nums[0]，如果大于nums[0]则说明nums[1]可以和nums[0]拼起来成为递增子序列。
        // 故nums[1]=2,否则为1
        dp[0] = 1;
        for(int i = 1;i<nums.length;++i){
            dp[i] = 1;
            // 计算nums[i] 是否可以和所有下标小于i的元素拼接成递增子序列。
            for(int j = 0;j<i;++j){
                if(nums[i]>nums[j]){
                    // 如果位置j的元素可以和位置i的元素拼成递增子序列则拼成的递增子序列长度为位置j的长度+1 即 dp[j]+1
                    // 又因为我们定义的是最长的递增子序列长度，我们遍历所有小于该元素的元素求出最大的，
                    // 因此每次都要和已有值比较，如果当前序列和前面的比比较大才替换进去。
                    dp[i] = max(dp[j]+1,dp[i]);
                }
            }
            // dp[i]计算完毕，保存全局最大值
            max_len = max(dp[i],max_len);
        }
        return max_len;

    }
}
