// 采用反转一半的思想
bool isPalindrome(int x) {
    // 去除负数, 因为负号的存在所以不会成为回文数
    if(x<0){return false;}
    // 对于尾数为0的数, 其最高位应为0(其实只有0符合), 故去除最低位为0, 且不等于0的数
    if((x%10 == 0 ) &&( x !=0 )){return false;}
    int reversed = 0;
    while(x>reversed){
        reversed *=10;
        reversed += x%10;
        x/=10;
    }
    if(reversed ==  x){
        return true;
    }else if (reversed > x){
        // 这里判断是否为位数为奇数的回文数(中间单独一个数字, eg. 121). 
        // 对于这样的数, 中间的那位数字会加在反转后的那部分数字上, 因此去除最低位后比较
        return x == (reversed/10);
    }else{
        return false;
    }
}
