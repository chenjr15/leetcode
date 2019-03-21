#define DEBUG 1
#include<stdio.h>
#include<limits.h>
int reverse(int x){
    long result=0;
    int flag = x<0?-1:1;
    char digits[32] = {0};
    unsigned ptr=0;
    int r=0;
    // 这里会溢出 因为正数范围比负数范围大1
    //x = x<0?-x:x;
    //去掉符号来算
    long y=x;
    y= y<0?-y:y;

    while(y){
        *(digits+ptr) =y%10;
        ptr++;
        y/=10;
    }
#if DEBUG     
        int i=ptr;
        putchar('#');
        while(i){
        putchar(*(digits+i) +'0' );
        --i;}
        putchar('\n');
#endif
    char* p = digits;
    while(ptr){
            result *=10;
            result += *(p++);
            --ptr;
    }
    result *= flag;
    //printf("%ld",result);
    if(result>INT_MAX || result<INT_MIN)
        result= 0;
    return (int)result;
}
int main(){
	int x = 900000;
	printf("%d\n",reverse(x));
	printf("%d\n",reverse(-x));
	x = 120450;
	printf("%d\n",reverse(x));
	printf("%d\n",reverse(-x));
	return 0;
}
