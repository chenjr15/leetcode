class Solution {
public:
    vector<int>* coins;
    vector<int>* table;
    int dp(int amount){
        //递归备忘录解法，自顶向下
        if (amount ==0){
            return 0;
        }
        if (amount<0){
            return -1;
        }
        if ((*this->table)[amount]!=-2 ){
            return (*this->table)[amount];
        }
        int res = amount+1;
        for(int coin:*coins){
            // 遍历每一种硬币凑法
            if(amount -coin <0){
                continue;
            }
            int sub = dp(amount-coin);
            // 子问题无解，跳过该子问题
            if (sub == -1 ){
                continue;
            }
            res = min(res,1+sub);
        }
        // 没有任何一个子问题有解，则该输入无解
        if (res == amount+1){
            res= -1;
        }
        (*this->table)[amount] = res;
        return res;
    }
    int coinChange(vector<int>& coins, int amount) {
        
        this->coins = &coins;
        this->table = new vector<int>(amount+1,-2) ;
        return dp(amount);
    }
};
