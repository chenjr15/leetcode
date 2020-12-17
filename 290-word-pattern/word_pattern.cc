class Solution {
public:
    bool cmp(const char * a,const char* b){
        
        while(*a&&*b&&*a==*b&&*a!=' '){
            // printf("%c==%c\n",*a,*b);
            ++a;++b;
        }
        return ((*b == 0|| *b==' ') && (*a ==0||*a==' ') );
    }
    bool wordPattern(string pattern, string s) {
        /**
        *  注意边界条件！！！
        */
        const char* table[26] = {0};
        const char* str = s.c_str();
        auto global_start = str;
        // 字符串的起始点
        auto start = str;

        bool match = true;
        for(auto c:pattern){
            // 1. pattern 比 str 长
            if(!*str){ match=false;break;};
            start = str;
            while(*str!=' '&& *str)++str;
            if(*str==' ') ++str;
            // printf("c=%c,start=%s,table=%s\n",c,start,table[c-'a']);
            if (!table[c-'a']){
                // 2. 判断是否和以前的pattern 相等，如果有相等的话是不对的，比如 ab -> dog dog 
                for(char c = 'a';c<'z'+1;++c){
                    if (table[c-'a'] && cmp(table[c-'a'],start)){
                        match = false;
                        break;
                    }
                }
                if(!match){
                    break;
                }
                table[c-'a'] = start;
            }else if (!cmp(table[c-'a'],start)){
                // pattern mismatch
                match = false;
                break;
            }
        }
        // printf("*last = %c\n",*str);
        // 3. str可能比pattern 更长
        return match && *str == 0;
       
    }
};
