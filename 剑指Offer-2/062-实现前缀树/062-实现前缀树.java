class Trie {
    Trie[] children;
    // int childrenCnt=0;
    boolean end=false;
    
    /** Initialize your data structure here. */
    public Trie() {
        children= new Trie[26];

    }
    
    /** Inserts a word into the trie. */
    public void insert(String word) {
        Trie p = this;
        int len = word.length();

        for(int i = 0;i<len;++i){
            int idx = word.charAt(i)-'a';
            if(p.children[idx] ==null){
                p.children[idx] = new Trie();
            }
            p = p.children[idx];
        }
        p.end = true;

    }
    
    /** Returns if the word is in the trie. */
    public boolean search(String word) {
        Trie p = this;
        int len = word.length();
        for(int i = 0;i<len;++i){
            int idx = word.charAt(i)-'a';
            if(p.children[idx] ==null){
                return false;
            }
            p = p.children[idx];
        }
        return p.end;
    }
    
    /** Returns if there is any word in the trie that starts with the given prefix. */
    public boolean startsWith(String word) {
        Trie p = this;
        int len = word.length();
        for(int i = 0;i<len;++i){
            int idx = word.charAt(i)-'a';

            if(p.children[idx] ==null){
                return false;
            }
            p = p.children[idx];
        }
        return true;

    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */