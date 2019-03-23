bool isPalindrome(int x) {
    if(x<0){return false;}
    if(x == 0){return true;}
    // 这个大概也是常数级的空间吧（大雾）
    char buffer[11] = {0};
    char* p = buffer;
    while(x){
        *p++ = x%10;
        x/=10;
    }
    
    int i,j;
    i=0;j=p-buffer-1;
    p = buffer;
    while(j>i){
        if(*(p+i) != *(p+j)){
            break;
        }
        i++;j--;
    }
    if(i >= j){
        return true;
    }else{
        return false;half
    }
}
