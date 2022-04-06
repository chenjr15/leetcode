# 剑指 Offer 37.序列化二叉树

难度：`Hard`

 链接：[https://leetcode-cn.com/problems/xu-lie-hua-er-cha-shu-lcof/](https://leetcode-cn.com/problems/xu-lie-hua-er-cha-shu-lcof/)

## 题目描述

<p>请实现两个函数，分别用来序列化和反序列化二叉树。</p>

<p>你需要设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。</p>

<p><strong>提示：</strong>输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅&nbsp;<a href="https://support.leetcode-cn.com/hc/kb/article/1194353/">LeetCode 序列化二叉树的格式</a>。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/09/15/serdeser.jpg" style="width: 442px; height: 324px;" />
<pre>
<strong>输入：</strong>root = [1,2,3,null,null,4,5]
<strong>输出：</strong>[1,2,3,null,null,4,5]
</pre>

<p>&nbsp;</p>

<p>注意：本题与主站 297 题相同：<a href="https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/">https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/</a></p>

## 标签

 - Tree 
 - Depth-First Search 
 - Breadth-First Search 
 - Design 
 - String 
 - Binary Tree 

## 代码

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Codec {

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        // 层次遍历
        if (root == null){
            return "";
        }
        ArrayList<TreeNode> queue = new ArrayList<>();
        queue.add(root);
        int head = 0;
        TreeNode node ;
        // lastNoneNull 用来截断后面连续的多余null
        int lastNoneNull = 1;
        for(;head<queue.size();++head){
            // 获取队头
            node =  queue.get(head);
            if (node ==null){
                continue;
            }
            lastNoneNull = head+1;
            queue.add(node.left);
            queue.add(node.right);
        }
        
        // 开始序列化

        StringBuilder sb = new StringBuilder();
        
        sb.append(root.val);
        for(head=1;head<lastNoneNull;++head){
            node =  queue.get(head);
            sb.append(',');
            if (node == null){
                sb.append('x');
            }else{            
                sb.append(node.val);
            }
        }
        // System.out.println("encoded:"+sb.toString());
        return sb.toString();
    }


    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        TreeNode root = null;
        // System.out.println("decoding:"+data);

        if (data==null || data==""){
            return root;
        }
        String[] nodeStrings = data.split(",");
        TreeNode[] nodeRef = new TreeNode[nodeStrings.length];
        TreeNode node;
        TreeNode parent;
        int p = 0;
        root =  new TreeNode(Integer.parseInt(nodeStrings[0]));
        nodeRef[0] = root;
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
    
        for(int i=1;i<nodeStrings.length;i+=2){
            parent = queue.poll();
            // 左节点
            if (!"x".equals(nodeStrings[i])){
                node =  new TreeNode(Integer.parseInt(nodeStrings[i]));
                parent.left = node;
                queue.offer(node);
            }
            
            if ((i+1)>=nodeStrings.length){
                break;
            }
            // 右节点
            if (!"x".equals(nodeStrings[i+1])){
                // 存在节点，
                node = new TreeNode(Integer.parseInt(nodeStrings[i+1]));
                parent.right = node;
                queue.offer(node);

            }
        }
        return root;
        
    }
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));
```