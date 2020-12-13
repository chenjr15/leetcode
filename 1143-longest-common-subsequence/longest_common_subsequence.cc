#define loop(i,init,stop) for(size_t (i) = (init);(i)<(stop);++i)
#define rloop(i,init,stop) for(size_t (i) = (init);(i)<(stop);--i)
class Solution {
public:
    int longestCommonSubsequence(string text1, string text2) {
        if (text2.size()==0 || text1.size()==0){
            return 0;
        }
    

        // 最长公共子序列
        // 如果 text1[n1-1] == text2[n2-1] ()
        // 则text1[n1-1] 一定在LCS里面,递归计算 text1[:n1] 和 text2[:n2] 
        // 如果他们不相等则，在text1[:n1]与text2的LCS 和 text1 与 text2[:n1] 中，选最大的那个

        // 自底向上算法
        // dp[i][j] 表示text1前i个字符 和 text2 前j个字符的最长上升子序列LCS
        // base case dp[*][0] = 0; dp[0][*] = 0;
        // 定义二维dp table
        vector<vector<size_t>> dp(text1.size()+1,vector<size_t>(text2.size()+1));

        loop(i,0,text1.size()) dp[i][0]=0;
        loop(i,0,text2.size()) dp[0][i]=0;
        loop(i,1,text1.size()+1){
            loop(j,1,text2.size()+1){
                if(text1[i-1] == text2[j-1]){
                    dp[i][j]=dp[i-1][j-1]+1;
                    continue;
                }
                dp[i][j] = max(dp[i-1][j],dp[i][j-1]);
            }
        }
        return dp[text1.size()][text2.size()];

    }
};
