class Trie {
    char c;
    Trie[]  children;
    boolean wordEnd = false;

    /** Initialize your data structure here. */
    public Trie() {
        this.children = new Trie[26];

    }
    public Trie(char c){
        this();
        this.c = c;
    }
    
    /** Inserts a word into the trie. */
    public void insert(String word) {
        Trie cur = this;
        int ci;

        for(int i = 0;i<word.length();i++){
            ci = word.charAt(i)-'a';
            if(cur.children[ci]==null){
                // 插入新节点
                cur.children[ci] = new Trie(word.charAt(i));
            }
            cur = cur.children[ci];
        }
        cur.wordEnd = true;

    }
    
    /** Returns if the word is in the trie. */
    public boolean search(String word) {
        Trie cur = this;
        int ci;

        for(int i = 0;i<word.length();i++){
            ci = word.charAt(i)-'a';
            if(cur.children[ci]==null){
                return false;
            }
            cur = cur.children[ci];
        }
        return cur.wordEnd;

    }
    
    /** Returns if there is any word in the trie that starts with the given prefix. */
    public boolean startsWith(String prefix) {
         Trie cur = this;
        int ci;

        for(int i = 0;i<prefix.length();i++){
            ci = prefix.charAt(i)-'a';
            if(cur.children[ci]==null){
                return false;
            }
            cur = cur.children[ci];
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
