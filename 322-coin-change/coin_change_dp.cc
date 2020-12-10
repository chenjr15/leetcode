class Solution {
public:

    int coinChange(vector<int>& coins, int amount) {
        // 自底向上的动态规划
        auto table =   vector<int>(amount+1,amount+1);
        // base case，所有状态从这个状态出发
        table[0] = 0;
        for (int n=1;n<=amount;++n){
            // 从1构建整个table
            for(auto coin :coins){
                // 状态转移，尝试选择所有可选的硬币，并选择用的硬币数最小的方案
                if(n-coin >=0){
                    table[n] = min(table[n],1+table[n-coin]);
                }
            }
        }
        if (table[amount]>amount){
            table[amount] = -1;
        }
        return table[amount];
    }
};
