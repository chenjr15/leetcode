     class Solution {
        public List<List<Integer>> findWinners(int[][] matches) {
            HashMap<Integer, Integer> winner = new HashMap<>();
            HashMap<Integer, Integer> loser = new HashMap<>();

            // int[][] result = new int[2][];
            for (int[] match : matches) {
                Integer winCnt = winner.getOrDefault(match[0], 0);
                Integer loseCnt = loser.getOrDefault(match[1], 0);
                winner.put(match[0], winCnt + 1);
                loser.put(match[1], loseCnt + 1);
            }
            ArrayList<List<Integer>> result = new ArrayList<>();
            // 计算结果
            ArrayList<Integer> ans0 = new ArrayList<>();
            for (Map.Entry<Integer, Integer> entry : winner.entrySet()) {
                Integer key = entry.getKey();
                Integer value = entry.getValue();
                if (loser.getOrDefault(key, 0).equals(0) )
                    ans0.add(key);
            }
            
            ans0.sort(Comparator.comparingInt(a -> a));
            result.add(ans0);
            ArrayList<Integer> ans1 = new ArrayList<>();
            for (Map.Entry<Integer, Integer> entry : loser.entrySet()) {
                Integer idx = entry.getKey();
                Integer times = entry.getValue();
                if (times == 1 )
                    ans1.add(idx);
            }
            ans1.sort(Comparator.comparingInt(a -> a));

            result.add(ans1);

            return result;


        }
    }