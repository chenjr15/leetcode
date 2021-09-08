class Solution {
    public int search(int[] nums, int target) {
        if(nums[0]>target || nums[nums.length-1]<target){
            return -1;
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
        return -1;

    }
}
