class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int total_size = nums1.length +nums2.length;
        int center =(total_size/2); 
        boolean odd = true;
        int[] selected = null;
        double median=0;
        int x=0,y=0;
        for(int i =0;i< center;i++){
            if(x<nums1.length && y<nums2.length){
                if(nums1[x]<nums2[y]){
                    x++;
                    selected=nums1;
                }else{
                    y++;
                    selected=nums2;
                }
            }else{
                if(x<nums1.length){
                    x++;
                    selected=nums1;

                }else{
                    y++;
                    selected=nums2;

                }
            }
        }
        // median = selected[((selected == nums1)?x:y)-1];
        if(x<nums1.length && y<nums2.length){
                if(nums1[x]<nums2[y]){
                    median += nums1[x];
                }else{
                    median += nums2[y];
                }
            }else{
                if(x<nums1.length){
                    median += nums1[x];

                }else{
                   median += nums2[y];

                }
            }
        int next;
        if(2*center == total_size){
             if(x<nums1.length && y<nums2.length){
                if(nums1[x]<nums2[y]){
                    median += nums1[x];
                }else{
                    median += nums2[y];
                }
            }else{
                if(x<nums1.length){
                    median += nums1[x];

                }else{
                   median += nums2[y];

                }
            }
            median /=2;
        }
        
        return median;
    }
    public static final int  smaller(int[] nums1, int[] nums2,Interger idx1,Interger idx2){
        
    }
}
