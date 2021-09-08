class Solution {
    public int searchInsert(int[] nums, int target) {
        // 同二分法，但是判断一下范围外的情况
        if(nums[0]>target ){
            return 0;
        }
        if(nums[nums.length-1]<target){
            return nums.length;
        }

        int l = 0;
        // 右边不减一即可，
        int r = nums.length;
        int center;
        while(l<r){
            center = (l+r) /2 ;
            if (nums[center] == target){
                return center;
            }
            if (target > nums[center]){
                // target 在右边，
                l = center+1;
            }else{
                // target 在左边
                r = center;
            }
        }
        // 直接返回左边即可
        return l;
    }
}
