class Solution {
        public int convertTime(String current, String correct) {
            int curTS = 0;
            int targetTS = 0;
            int  total = 24*60;

            String[] parts = current.split(":");
            curTS = Integer.parseInt(parts[0]) * 60 + Integer.parseInt(parts[1]);
            parts = correct.split(":");
            targetTS = Integer.parseInt(parts[0]) * 60 + Integer.parseInt(parts[1])  ;
            int diff = targetTS-curTS;
            // diff%=total;
            int cnt = 0 ;
            while (diff>0){
                if(diff>=60){
                    cnt += diff/60;
                    diff%=60;
                }else if (diff>=15){
                    cnt += diff/15;
                    diff%=15;
                }else if (diff>=5){
                    cnt += diff/5;
                    diff%=5;
                }else {
                    cnt+=diff;
                    diff=0;
                }
            }

            return cnt;

        }
}