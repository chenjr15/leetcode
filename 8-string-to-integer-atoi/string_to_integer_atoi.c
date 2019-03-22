#include<string.h>
#include<ctype.h>

int myAtoi(char* str) {
    if(!str) return 0;
    unsigned long sum=0;
    int result=0;
    // 跳过空白字符
    while(*str == ' ')str++;
    // 将负号存起来单独放,简化计算
    int minus_flag = 0;
    if(*str == '+'){
        str++;
    }else if (*str == '-'){
        str++;
        minus_flag=1;
    }
    // 遇到非数字就跳出循环
    while(isdigit(*str)){
        sum *= 10;
        sum += ( *str - '0');
        str++;
        if(sum>((long)INT_MAX+minus_flag)){
            sum =minus_flag?INT_MIN:INT_MAX;
            break;
        }
    }
    result = minus_flag?(-sum):sum;
    return result;
}
