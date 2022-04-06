class TireTree{
    // char chr;
    // 当前节点包含字符
    String str;
    // TireTree parent;
    TireTree[] children;
    public TireTree(){
        this.children = new TireTree[26];
    }
    // public TireTree(char chr,TireTree parent ){
    //     this();
    //     this.chr = chr;
    //     this.parent = parent;
        
    // }
    public TireTree(String str){
        // this(chr,parent);
        this.str  = str;
    }
    /**
    * Search 根据给定字符串路径搜索到指定的节点，并不会修改内容，树中没有该字符串也会返回对应节点(如果存在)
    * 由于题目的特殊性，直接判断是不是每个节点都有存值，即p.str == null
    *
     */
    public TireTree Search(String str){
        if (str==null||str==""){
            return this;
        }
        TireTree p = this.children[str.charAt(0)-'a'];
        for(int i = 1 ;i<str.length();++i){
            // 逐个字符查找
            if (p==null || p.str == null ){
                return null;
            }
            p=p.children[str.charAt(i)-'a'];
        }
        // 这里不判断p.str == str
        return p;
    }

    public TireTree Insert(String str){
        if (str==null||str==""){
            return this;
        }
        TireTree p = this;
        
        for(int i = 0 ;i<str.length();++i){
            // 逐个字符插入
            if (p.children[str.charAt(i)-'a'] == null){
                // 创建新的节点
                p.children[str.charAt(i)-'a'] = new TireTree();
            }
            p=p.children[str.charAt(i)-'a'];
        }
        // 给节点赋值
        p.str = str;
        return p;
    }
}
class Solution {
    public String longestWord(String[] words) {
        TireTree root = new TireTree();
        // 字典树
        // 题目有坑，要求逐个字符加上来到底，即ewqz 要有e，ew，ewq，三个在里面，因此直接在search的时候判断死不是每个节点都有就行了。
        String longest = "";
        TireTree p;
        // ArrayList<TireTree> 
        // 构建字典树
        for(String s:words){

            // 逐个插入字典树
            p = root.Insert(s);
            
        }
        for(String s:words){
            if(s==null||s==""){
                continue;
            }
            // 逐个插入字典树
            p = root.Search(s);
            if (p!=null &&
             (s.length()>longest.length() || s.length() == longest.length()&& longest.compareTo(s)>0)){
                longest = s;
            } 
        }

        return longest;
    }
}