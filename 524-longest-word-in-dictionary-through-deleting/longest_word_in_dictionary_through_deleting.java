class Solution {
    public boolean match(String key,String  str){
        int strlen = str.length();
        int keylen = key.length();
        if(keylen>strlen){
            return false;
        }
        int i=0;
        int j=0; 
        
        while((strlen-j)>=(keylen-i) && i<keylen && j<strlen){
            
            if(key.charAt(i) == str.charAt(j) ){
                ++i;
            }
            ++j;
            
        }
    
        // 条件是j走的长度大于str的长度且i走完key
        return j>=keylen && i==keylen;
        
        
    }
    public String findLongestWord(String s, List<String> dictionary) {
        // 首先要满足可匹配，然后长度要最长且字典序最小

        // 先对字符串字典排序

        Collections.sort(dictionary,(a,b)->{
                // 按长度逆序
                if(a.length()!=b.length())
                    return b.length()-a.length();
                // 字典顺序
                return a.compareTo(b);
        });
        String ret = "";
        
        for(String d:dictionary){
            if(match(d,s)){
                ret = d;
                break;
            }
        }
        return ret;
        

    }
}
