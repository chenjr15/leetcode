class Solution {
    public long[] kthPalindrome(int[] queries, int intLength) {
        long[] answer = new long[queries.length];
        // 中间的位数， 取值范围是0-9 共十位，
        long centerLen = (intLength - 1) / 2 ;
        // 个数最多为，头尾的取值范围是1-9 共9位，乘上 中间的位数，每位有10个选项
        long maxCnt = (long) (9 * Math.pow(10, centerLen));
        byte [] answerParts  = new byte[intLength];

        for (int i = 0, queriesLength = queries.length; i < queriesLength; i++) {
            int q = queries[i];
            // 计算每个查询
            if (q > maxCnt) {
                answer[i] = -1;
            }else{
                q--;
                // 先计算一半, 计算右边
                int right = intLength /2 ;
                while(right<intLength-1){
                    answerParts[right] = (byte) (q%10);
                    q/=10;
                    right++;
                }
                answerParts[intLength-1]= (byte) (q+1);

                //
                for (int l =0,r=intLength-1;l<r;l++,r--){
                    answerParts[l] = answerParts[r];
                }
                // 转换成long
                answer[i]= 0;
                for (byte part : answerParts) {
                    answer[i]*=10;
                    answer[i]+=part;
                }
                // System.out.println("q = "+(queries[i])+" "+Arrays.toString(answerParts)+answer[i]);
            }
        }
        return answer;
    }
}