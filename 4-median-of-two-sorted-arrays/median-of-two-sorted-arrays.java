class Solution {
    private int idx1=0,idx2=0;
    private int[] nums1=null;
    private int[] nums2 = null;
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        this.nums1 = nums1;
        this.nums2 = nums2;
        int total_size = nums1.length +nums2.length;
        int center =(total_size/2); 
        double median=0;
        if((total_size - 2*center)==0){
            // 少一次
            median = this.fastfoward(center-1);
            median = (median+fastfoward(1))/2;
        }else{
            // center次
            median = fastfoward(center);
        }
        
        
        return median;
    }
    public int  fastfoward(int count){
        int len1=nums1.length;
        int len2 = nums2.length;
        while((count--) != 0){
            if(idx1<len1&& idx2<len2){
                if(nums1[idx1]<nums2[idx2]){
                    idx1++;
                }else{
                    idx2++;
                }
            }else{
                if (idx1<len1){
                    idx1++;
                }else{
                    idx2++;
                }
            }
        }
        int smaller = 0;
        // 返回快进指针指向的比较小的元素
        if(idx1<len1&& idx2<len2){
                if(nums1[idx1]<nums2[idx2]){
                    smaller = nums1[idx1]; 
                }else{
                    smaller = nums2[idx2]; 
                }
            }else{
                if (idx1<len1){
                    smaller = nums1[idx1]; 
                }else{
                    smaller = nums2[idx2]; 
                }
            }
        return smaller;
    }
}
