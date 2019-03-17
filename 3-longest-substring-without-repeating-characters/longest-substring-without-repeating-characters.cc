
#include <string.h>
using namespace std;
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        unsigned longest = 0;
        unsigned got_size=0;
        int buf[256]={0};
        int*  got = &buf[0];
        memset(got,-1,sizeof(int)*255);
        for(int i =0;i<s.size();i++){
            if ( got[s[i]]!=-1) {
                longest = longest>got_size?longest:got_size;
                i = got[s[i]];
                memset(got,-1,sizeof(int)*256);
                got_size=0;
            } else {
                got[s[i]] = i;
                got_size++;
            }
        }
        return  longest>got_size?longest:got_size;
        
    }
};
