class Solution {
public:
    void sortColors(vector<int>& nums) {
        unsigned low = 0;
        unsigned p = 0;
        int t;
        for(;p<nums.size();++p){
            if (nums[p] == 0){
                t = nums[p];
                nums[p] = nums[low];
                nums[low] = t;
                ++low;
            }
        }
        for(p = low;p<nums.size();++p){
            if (nums[p] == 1){
                t = nums[p];
                nums[p] = nums[low];
                nums[low] = t;
                ++low;
            }
        }
    }
};
