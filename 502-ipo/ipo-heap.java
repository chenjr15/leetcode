class Solution {
    public int findMaximizedCapital(int k, int w, int[] profits, int[] capital) {

        int n = profits.length;
        // 贪心，每次在可以投资的项目中选择利润最高的
        // 1. 基于capital排序，在所有可以选择的capital里寻找最大的profit即可


        // 最多投资的数目一定是小于项目数和限制数的
        int maxTimes = n>k?k:n;

        // 合并到一个数组中
        int[][] projects = new int[n][2];
        for (int i =0;i<n;++i){
            projects[i][1] = capital[i];
            projects[i][1] = profits[i];
        }
        // 基于capital排序
        Arrays.sort(projects, (a, b) -> a[0] - b[0]);
        // 遍历寻找可投资项目最大利润的做法会超时，只能用堆
        // 堆中维护的是所有可以投资的项目的profit,
        // 每次拿一个最大的profit，然后更新w，将所有可以投资的profit加入堆中

        int i = 0;
        PriorityQueue<Integer> maxHeap = new PriorityQueue<>((x, y) -> y - x);
        for(int times =0;times<maxTimes;++times){
            // 遍历寻找可以投资的项目
            while(i<n&&projects[i][0]<=w){
                // 所有现在可以投资的项目profit加入堆中
                maxHeap.add(projects[i][1]);
                ++i;
            }
            if(maxHeap.isEmpty()){
                break;
            }
            w += maxHeap.poll();
        }
        return w;
    }
}
