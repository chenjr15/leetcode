int search(int* nums, int numsSize, int target) {
    while(numsSize){
        if(nums[--numsSize]==target){
            break;
        }
    }
    return numsSize;
}
