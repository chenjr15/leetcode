class Solution {
    public int findMaximizedCapital(int k, int w, int[] profits, int[] capital) {

        int n = profits.length;
        // 贪心，每次在可以投资的项目中选择利润最高的
        // 1. 基于capital排序，在所有可以选择的capital里寻找最大的profit即可

        // 在有序的情况下
        // 最多投资的数目一定是小于项目数和限制数的
        int maxTimes = n>k?k:n;

        // 合并到一个数组中
        int[][] projects = new int[n][2];
        for (int i =0;i<n;++i){
            projects[i][0] = capital[i];
            projects[i][1] = profits[i];
        }
        // 基于capital排序
        Arrays.sort(projects, (a, b) -> a[0] - b[0]);
        for(int times =0;times<maxTimes;++times){
            // 遍历寻找可以投资的项目
            int maxI = -1;
            for(int i =0 ;i<n;++i){
                if(projects[i][0] <0){
                    continue;
                }               
                if(projects[i][0]>w){
                    break;
                }
                // 所有可以选择且未被选择的项目
                if(maxI==-1||(projects[i][1]>projects[maxI][1])){
                    maxI = i;
                }
            }
            if(maxI>-1){
                w += projects[maxI][1];
                // 设置标志位
                projects[maxI][0] = -1;
            }
        }
        return w;
    }
}

