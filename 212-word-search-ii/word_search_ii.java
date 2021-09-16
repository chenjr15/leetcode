class Solution {
    Trie tree;
    String[] words;
    char[][] board;
    List<String> result;
    boolean[][] visited;
    public List<String> findWords(char[][] board, String[] words) {

        // 将words逐个加入前缀树中
        tree = new Trie();
        for(String word :words){
            tree.insert(word);
        }
        this.board = board;
        this.words = words;
        result = new ArrayList<String>(words.length);
        this.visited = new boolean[board.length][board[0].length];

        for(int i=0;i<board.length;++i){
            for(int j=0;j<board[0].length;++j){
                // 遍历每个起点
                // 排除不为给定单词首字母的位置
                if(!tree.charIn(board[i][j])){
                    continue;
                }
                traceBack(i,j,0);
                // 数量达到后不要重复寻找
                if(result.size()>= words.length){
                    break;
                }

            }
        }
        return result;

    }
    public void traceBack(int i,int j,int lastDir){
       
        // lastDir
        // 1 上
        // 2 下
        // 3 左
        // 4 右
        if(i<0||i>=board.length){
            return;
        }
        if(j<0||j>=board[0].length){
            return;
        }
        if (visited[i][j]){
            return;
        }
        Trie cur=  tree.stepIn(board[i][j]);
        if(cur!=null){
            if(cur.word!=null){
                this.result.add(cur.word);
                // 单词匹配到要记得删掉，防止重复匹配
                cur.word=null;
            }
            visited[i][j] = true;
            // 向上
            if(lastDir!=2 )
                traceBack(i-1,j,1);
            // 向下
            if(lastDir!=1)
                traceBack(i+1,j,2);
            // 向右
            if(lastDir!=3)
                traceBack(i,j+1,4);
            // 向左
            if(lastDir!=4)
                traceBack(i,j-1,3);
            visited[i][j]=false;
            tree.stepOut();
        }

    }

}


class Trie {
    char c;
    Trie[]  children;
    String word;
    List<Trie> path;
    Trie cur;

    /** Initialize your data structure here. */
    public Trie() {
        this.children = new Trie[26];
        this.path = new ArrayList<Trie>();
        this.cur = this;

    }
    public Trie(char c){
        this();
        this.c = c;
    }

    public boolean charIn(char c){
        return this.children[c-'a']!=null;
    }

    public Trie stepIn(char c){
        Trie next = this.cur.children[c-'a'];
        if(next!=null){
            this.path.add(next);
            this.cur = next;
        }
        return next;
    }
    public void stepOut(){
        if(this.path.size()==0){
            return;
        }
        if(this.path.size()>1){
            this.path.remove(this.path.size()-1);
            this.cur = this.path.get(this.path.size()-1);
        }else{
            this.path.remove(0);
            this.cur = this;
        }
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
        cur.word = word;

    }
    
    public String toString(){
        return  ""+this.c;
    }

    // /** Returns if the word is in the trie. */
    // public boolean search(String word) {
    //     Trie cur = this;
    //     int ci;

    //     for(int i = 0;i<word.length();i++){
    //         ci = word.charAt(i)-'a';
    //         if(cur.children[ci]==null){
    //             return false;
    //         }
    //         cur = cur.children[ci];
    //     }
    //     return cur.wordEnd;

    // }

    // /** Returns if there is any word in the trie that starts with the given prefix. */
    // public boolean startsWith(String prefix) {
    //      Trie cur = this;
    //     int ci;

    //     for(int i = 0;i<prefix.length();i++){
    //         ci = prefix.charAt(i)-'a';
    //         if(cur.children[ci]==null){
    //             return false;
    //         }
    //         cur = cur.children[ci];
    //     }
    //     return true;

    // }
}

