class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        // 子数组一定是连续的
        // 定义dp[i] 为nums[i]结尾的最大子数组和，dp[0] = nums[0];
        // 对于每一个数只有选或者不选
        // 如果选，则dp[i] = dp[i-1]+nums[i];
        // 如果不选，则子数组中断，dp[i] = nums[i];
        // 如果nums[i] 即当前数是负数，肯定不能选，因为负数会让累加的值变小，不符合最大的定义
        // 如果是正数则选入
        // 因为每次只用到上次的值，所以dp[] 可以优化成一个上次的变量
        int dp = nums[0];
        int sum=dp;
        auto len =nums.size();
        for(size_t i=1;i<len;++i){
            int &n = nums[i];
            // if(n>0){
            //     dp +=n;
            // }else{
            //     dp = n;
            // }
            // auto chose = dp+n;
            // if (n>chose){
            //     dp = chose
            // }else{
            //     dp = n;
            // }
            // dp = (dp+n)>=dp?dp+n:n;
            dp = max(n,dp+n);
            
            // 从 1 到 n-1 选一个最大的结尾即可
            if (dp>sum){
                sum = dp;
            }
            // sum = max(dp,sum);

        }
        return sum;
    }
};
