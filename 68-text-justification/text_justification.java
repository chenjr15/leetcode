class Solution {
    public List<String> fullJustify(String[] words, int maxWidth) {
        // 并不难，但是细节比较麻烦
        List<String> ret = new ArrayList<String>();
        int i =0;
        // 使用StringBuilder避免频繁生成字符串
        StringBuilder builder = new StringBuilder(maxWidth);
        int wc = 0;
        int charUsed= 0;
        
        while(i<words.length){
            builder.setLength(0);
            // add first word.
            charUsed = words[i].length();
            wc=1;
            // 断行，检查下一个单词是否可以被加入该行
            while( 
                ((wc+i)<words.length) && 
                ((charUsed+words[i+wc].length() + wc) <= maxWidth)
            ){
                charUsed += words[i+wc].length();
                ++wc;
            }
           
            if((wc+i)>=words.length || wc==1){
                //判断最后一行
                for(int k = 0 ;k<wc-1;++k ){
                    builder.append(words[i+k]);
                    builder.append(' ');
                }
                // 最后一个不加空格
                builder.append(words[i+wc-1]);
                // 补齐空格
                while(builder.length()<maxWidth){
                    builder.append(' ');
                }
            }else{
                // 空格长度应该是maxWidth-所有单词长度然后平分
                int wordWidth = (maxWidth - charUsed)/(wc-1);
                // System.out.println("ww:"+wordWidth);
                int extraWidthIndex = (maxWidth - charUsed) - wordWidth*(wc-1);
                for(int k=0;k<wc-1;++k){
                    String word = words[i+k];
                    builder.append(word);
                    int spaceLeft = wordWidth + ((k<extraWidthIndex)?1:0);
                    // System.out.println(word+" "+spaceLeft);
                    // 插入空格
                    for(int s = 0;s<spaceLeft;++s )
                        builder.append(' ');
                }
                // 最后一个单词不加空格
                String word = words[i+wc-1];
                builder.append(word);
            }
            ret.add(builder.toString());
            i=i+wc;
        }
        return ret;
    }
}
