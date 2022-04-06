class Solution {
    public int minCostClimbingStairs(int[] cost) {
        if (cost.length <2){
            return -1;
        }
        // 每次跳一步或者两步，
        // dp 数组，表达跳到这一步需要的最小花费，那么跳到楼顶就是 min(dp[n-1]+cost[n-1],dp[n-2]+cost[n-2])
        int cost1,cost2,minCost;

        cost1 = cost[0];
        cost2 = cost[1];
        for (int i = 2;i<cost.length;i++){
            minCost = Math.min(cost1,cost2)+cost[i];
            cost1 = cost2;
            cost2 = minCost;
        }
        // System.out.println(Arrays.toString(minCost));
        return Math.min(cost2,cost1);
        

    }
}