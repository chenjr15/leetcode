int findMin(int* nums, int numsSize) {
    int i=0;
    int min=*nums;
    while(i<numsSize){
        if(nums[i]<min){
            min = nums[i];
        }
        i++;
    }
    return min;
    
}
