# 648.单词替换

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/replace-words/](https://leetcode-cn.com/problems/replace-words/)

## 题目描述

<p>在英语中，我们有一个叫做&nbsp;<code>词根</code>(root) 的概念，可以词根<strong>后面</strong>添加其他一些词组成另一个较长的单词——我们称这个词为&nbsp;<code>继承词</code>(successor)。例如，词根<code>an</code>，跟随着单词&nbsp;<code>other</code>(其他)，可以形成新的单词&nbsp;<code>another</code>(另一个)。</p>

<p>现在，给定一个由许多<strong>词根</strong>组成的词典 <code>dictionary</code> 和一个用空格分隔单词形成的句子 <code>sentence</code>。你需要将句子中的所有<strong>继承词</strong>用<strong>词根</strong>替换掉。如果<strong>继承词</strong>有许多可以形成它的<strong>词根</strong>，则用<strong>最短</strong>的词根替换它。</p>

<p>你需要输出替换之后的句子。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
<strong>输出：</strong>"the cat was rat by the bat"
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
<strong>输出：</strong>"a a b c"
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= dictionary.length&nbsp;&lt;= 1000</code></li>
	<li><code>1 &lt;= dictionary[i].length &lt;= 100</code></li>
	<li><code>dictionary[i]</code>&nbsp;仅由小写字母组成。</li>
	<li><code>1 &lt;= sentence.length &lt;= 10^6</code></li>
	<li><code>sentence</code>&nbsp;仅由小写字母和空格组成。</li>
	<li><code>sentence</code> 中单词的总量在范围 <code>[1, 1000]</code> 内。</li>
	<li><code>sentence</code> 中每个单词的长度在范围 <code>[1, 1000]</code> 内。</li>
	<li><code>sentence</code> 中单词之间由一个空格隔开。</li>
	<li><code>sentence</code>&nbsp;没有前导或尾随空格。</li>
</ul>

<p>&nbsp;</p>

## 标签

 - Trie 
 - Array 
 - Hash Table 
 - String 

## 代码

```java
class Solution {
    public String replaceWords(List<String> dictionary, String sentence) {
        Trie dict = new Trie();
        for (String s : dictionary) {
            dict.insert(s);
        }

        // 遍历单词
        String[] words = sentence.split(" ");
        StringBuilder sb = new StringBuilder();
        for (String word : words) {
            String prefix = dict.findPrefix(word);
            if (prefix == null) {
                prefix = word;
            }
            sb.append(prefix);
            sb.append(' ');
        }
        if (sb.length() > 1)
            sb.deleteCharAt(sb.length() - 1);
        return sb.toString();
    }
}

class Trie {
    Trie[] children;
    // int childrenCnt=0;
    boolean end = false;
    String s;

    /**
     * Initialize your data structure here.
     */
    public Trie() {
        children = new Trie[26];

    }

    /**
     * Inserts a word into the trie.
     */
    public void insert(String word) {
        Trie p = this;
        int len = word.length();

        for (int i = 0; i < len; ++i) {
            int idx = word.charAt(i) - 'a';
            if (p.children[idx] == null) {
                p.children[idx] = new Trie();
            }
            p = p.children[idx];
        }
        p.end = true;
        p.s = word;

    }

    /**
     * Returns if the word has a prefix in the trie.
     */
    public String findPrefix(String word) {
        Trie p = this;
        int len = word.length();
        for (int i = 0; i < len; ++i) {
            if (p.s != null) {
                break;
            }
            int idx = word.charAt(i) - 'a';
            if (p.children[idx] == null) {
                break;
            }
            p = p.children[idx];
        }
        return p.s;
    }

    /**
     * Returns if the word is in the trie.
     */
    public boolean search(String word) {
        Trie p = this;
        int len = word.length();
        for (int i = 0; i < len; ++i) {
            int idx = word.charAt(i) - 'a';
            if (p.children[idx] == null) {
                return false;
            }
            p = p.children[idx];
        }
        return p.end;
    }

    /**
     * Returns if there is any word in the trie that starts with the given prefix.
     */
    public boolean startsWith(String word) {
        Trie p = this;
        int len = word.length();
        for (int i = 0; i < len; ++i) {
            int idx = word.charAt(i) - 'a';

            if (p.children[idx] == null) {
                return false;
            }
            p = p.children[idx];
        }
        return true;

    }
}

```