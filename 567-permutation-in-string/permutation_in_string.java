class Solution {

    public boolean checkInclusion(String s1, String s2) {
        
        int len1 = s1.length();
        int len2 = s2.length();
        if(len1>len2){
            return false;
        }
        boolean contains = false;
      
        int[] s1map = new int[26];
        int[] s2map = new int[26];
        // 计算s1和s2和s1长度相等的字符出现频率
        char c;
        
        for(int i=0;i<len1;++i){
            ++s1map[s1.charAt(i)-'a'];
            ++s2map[s2.charAt(i)-'a'];
        }
        if(Arrays.equals(s1map,s2map)){
            return true;
        }
        for(int offset=len1;offset<len2;++offset){
            // 遍历每个长度为len1的窗口,[offset,len+offset)
            // 仅仅计算头尾的字符即可
            c = s2.charAt(offset);
            ++s2map[c-'a'];
            --s2map[s2.charAt(offset - len1)-'a'];
            if(Arrays.equals(s1map,s2map)){
                contains = true;
                break;
            }
        }
        return contains;

    }
}
