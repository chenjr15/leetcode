class KthLargest {

    Queue<Integer> minHeap;
    int k =0 ;

    public KthLargest(int k, int[] nums) {
        // 维护一个小根堆，里面保存的是最大的k个数，那么堆顶是这里面最小的，也就是第k大的数 。
        // 然后一往里加元素。加满了就把顶上去掉，如果是堆顶的元素大的话就把堆顶的去掉，插入新的。
        this.minHeap = new PriorityQueue<>();
        this.k = k ;
        for (int i =0;i<nums.length;i++){
            this.add(nums[i]);
        }

    }
    
    public int add(int val) {

        
        if(minHeap.size()<k){
            minHeap.add(val);
        }else if (minHeap.size()>=k && minHeap.peek() < val){
            // 去掉堆顶
            minHeap.poll();
            minHeap.add(val);
        }
        // System.out.println(minHeap.toString());
        return minHeap.peek();
    }
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * KthLargest obj = new KthLargest(k, nums);
 * int param_1 = obj.add(val);
 */