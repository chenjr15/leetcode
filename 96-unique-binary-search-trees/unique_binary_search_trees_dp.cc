class Solution {
public:
    int numTrees(int n) {
        if (n<0){
            return 1;
        }
        if (n<3){
            return n;
        }
        vector<int> nums(n+1,0);
        nums[0] = 1;
        nums[1]=1;
        nums[2] = 2;
        // 自底向上构建 动态规划
        for(int j = 3;j<=n;++j){
            // 遍历根节点所在的位置
            for(int i = 0;i<j;++i){
                // 左边的可能性 x 右边的可能性
                nums[j] += nums[i]*nums[j-i-1];
            }
        }
        return nums[n];
    }
};
