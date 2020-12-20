/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/

class Solution {
public:
    Node* l=NULL;
    Node* r = NULL;
    void dfs(Node * n){
        if(!n)return ;
        if(!n->left){
            // 底下没有子节点了，返回
            return ;
        }
        // 底下还有子节点
        l = n->left;
        r =  n->right;
        // 逐层向下遍历 左子树的右节点，右子树的左节点
        while(l){
            l->next = r;
            l = l->right;
            r = r->left;
        }
        dfs(n->left);
        dfs(n->right);

    }
    Node* connect(Node* root) {
        dfs(root);
        return root;
        
    }
};
