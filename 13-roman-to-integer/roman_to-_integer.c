int romanToInt(char* s) {
    int r=0;
    char next=0;
    while(*s){
        next=*(s+1);
        switch(*s){a
            case 0:
                break;
            case 'I':
                if(next=='V'||next=='X'){
                    r+=next=='V'?4:9;
                    s++;
                }else{
                    r+=1;
                }
                break;
            
            case 'X':
                if(next=='L'||next=='C'){
                    r+=next=='L'?40:90;
                    s++;
                }else{
                    r+=10;
                }
                break;
                
            case 'C':
                if(next=='D'||next=='M'){
                    r+=next=='D'?400:900;
                    s++;
                }else{
                    r+=100;
                }
                break;
                
            case 'V':
                r+=5;                
                break;
            case 'L':
                r+=50;                
                break;
            case 'D':
                r+=500;                
                break;
            case 'M':
                r+=1000;                
                break;
        }
        s++;
    }
    return r;
    
}
