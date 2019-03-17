class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
    int i,j;
    vector<int> result = {0,0};
    result[0]=0;
    result[1]=0;
    for(i=0;i<nums.size();i++){
        for (j = i+1;j<nums.size();j++){
            if((nums[i]+nums[j]) == target){
                result[0]=i;
                result[1]=j;
                break;
            }
        }
    }
    return result;
    }
};
