int compress(char* chars, int charsSize) {
    int p=0;
    int lastChar = 0;
    int newChars = 0;
    int repeatTimes = 0;
    char n_buf[5] = {0};
    int i =0;
    while(p<charsSize-1){
        if(chars[p] ==chars[p+1]){
            lastChar = p;
            while( p<charsSize){
                if (chars[p] == chars[lastChar])
                    ++p;
                else 
                    break;
            }
            repeatTimes = (p-lastChar);
            if(repeatTimes>1 ){
                i = 0;
                while(repeatTimes){
                    *(n_buf+i)= repeatTimes%10;
                    repeatTimes/=10;
                    ++i;
                }
                chars[newChars++] = chars[lastChar];
                while(i--){
                    chars[newChars++]= *(n_buf+i)+'0';
                }
            }else{
                newChars = p;
            }
            
            
        }else{
            chars[newChars++] = chars[p++];
            // p++;
        }
    }
    if(charsSize!=p)chars[newChars++] = chars[p++];
    return newChars;
    
}