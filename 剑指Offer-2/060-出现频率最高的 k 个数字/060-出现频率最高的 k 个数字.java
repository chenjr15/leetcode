class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        // hashmap + 最小堆或者快排一部分
        HashMap<Integer,Integer> map = new HashMap<>();
        for(int n:nums){
            Integer cnt = map.get(n);
            if (cnt==null){
                cnt=1;
            }else{
                cnt++;
            }
            map.put(n,cnt);
        }
        // 默认小顶堆，即每次弹出最小的
        PriorityQueue<Map.Entry<Integer,Integer>> pq = new PriorityQueue<>((e1,e2) -> e1.getValue()- e2.getValue());
        for (Map.Entry < Integer, Integer > entry: map.entrySet()) {
//            System.out.println(pq.toString());
            if(pq.size()<k){
                pq.offer(entry);
            }else{
                if (entry.getValue()>pq.peek().getValue()){
                    pq.poll();
                    pq.offer(entry);
                }
            }
        }
        int minlen = Math.min(k,pq.size());
        int[] result = new int[minlen];

        for (int i = 0; i < minlen; i++) {
            result[i] = pq.poll().getKey();
        }
        return result;

    }
}