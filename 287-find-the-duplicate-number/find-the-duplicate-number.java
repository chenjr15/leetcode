class Solution {
    public int findDuplicate(int[] nums) {
        int fast =nums[0];
        int slow = nums[0];
        // 走两步
            fast=nums[fast];
            fast=nums[fast];
            // 走一步
            slow=nums[slow];
        while(fast!=slow){
            // 走两步
            fast=nums[fast];
            fast=nums[fast];
            // 走一步
            slow=nums[slow];
        }
        
        fast = nums[0];
        while(fast!=slow){
            // 走一步
            fast=nums[fast];
            // 走一步
            slow=nums[slow];
            
        }
        return fast;
        
    }
}
