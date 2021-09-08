class Solution {
    public int[] plusOne(int[] digits) {
        // 需要处理进位
        int len =  digits.length-1;
        boolean remain = true;
        for(;len>=0;--len ){
            if(remain){
                ++digits[len];
                remain = false;
            }
            
            if(digits[len]==10){
                digits[len]=0;
                remain = true;
            }
        }
        if(remain){
            int[] digits_new  = new int[digits.length+1];
            for(len =0;len<digits.length;++len){
                digits_new[len+1] = digits[len];
            }
            digits = digits_new;
            digits[0] = 1;
        }
        return digits;

    }
}
