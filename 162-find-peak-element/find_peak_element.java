class Solution {
    public int findPeakElement(int[] nums) {
        
        int len = nums.length;
        if(len==1 || nums[0]>nums[1]){
            return 0;
        }
        if(nums[len-1]>nums[len-2]){
            return len-1;
        }
        int mid=len/2;
        int l=0;
        int r=len;
        while(l<r){
            mid = (l+r)/2;
            if(nums[mid] > nums[mid-1] && nums[mid]>nums[mid+1]){
                return mid;
            }
            if(nums[mid]>=nums[mid-1] && nums[mid]<=nums[mid+1]){
                // /-/  上升型，往右走
                l = mid+1;
            }else {
                // 下降型(或者是平的) ，往左走
                r = mid;
            }
        }
        return mid;
    }
}
