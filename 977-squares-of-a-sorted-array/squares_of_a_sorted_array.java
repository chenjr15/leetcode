class Solution {
    public int[] sortedSquares(int[] nums) {
        // 输入中包含负数且要求按结果有序，因此要按绝对值来计算。
        // 具体来说是按绝对值倒序计算，从两边往中间计算即可，结果从result的后面往前填入
            int[] result = new int[nums.length];
            int left=0,right = nums.length-1;
            int idx = right;
            while(left<right){
                if(Math.abs(nums[left])>Math.abs(nums[right])){
                    result[idx] = nums[left] *nums[left];
                    ++left;
                }else{
                    result[idx] = nums[right] *nums[right];
                    --right;
                }
                --idx;
            }
           
            result[idx] = nums[left]*nums[left];
           
            return result;
            
    }
}
