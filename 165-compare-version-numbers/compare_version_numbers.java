class Solution {
    public int compareVersion(String version1, String version2) {
        // 按.分割的字符串，每次比较一个版本号即可
        // 算是双指针的线性算法
        int ver1,ver2;
        int idx1=0,idx2=0;
        char c;
        while(idx1<version1.length() || idx2<version2.length()){
            ver1 = 0;
            // 获取版本1的修订号
            for(;idx1<version1.length();++idx1){
                c = version1.charAt(idx1);
                if(c == '.'){
                    ++idx1;
                    break;
                }
                ver1*=10;
                ver1+=c - '0';
            }
            
            ver2= 0;
            // 获取版本1的修订号
            for(;idx2<version2.length();++idx2){
                c = version2.charAt(idx2);
                if(c == '.'){
                    ++idx2;
                    break;
                }
                ver2*=10;
                ver2+=c - '0';
            }
            // 比较版本号若是不相等则可以可以直接返回结果，否则继续比较后面的版本号
            if (ver1!=ver2){
                return ver1>ver2?1:-1;
            }
        
        }
        return 0;
    }
}
