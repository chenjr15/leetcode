class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
        
        vector<int> dp(nums.size()+1,1);
        int max_result=1;
        dp[0] = 1;
        for(int i =1 ;i<nums.size();++i){
            for(int j =0;j<i;++j){
                // 遍历所有第i个数结尾的，能和第i个数构成递增子序列的
                if(nums[j]<nums[i]){
                    dp[i] = max(dp[j]+1,dp[i]);
                }
            }
            max_result = max(max_result,dp[i]);
        
        }
        return max_result;

    }
};
