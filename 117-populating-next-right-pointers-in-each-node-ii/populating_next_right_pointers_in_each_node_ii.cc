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

    Node* connect(Node* root) {
        if(!root){
            return root;
        }

        // 建立这一层的连接
        auto p = root;
        Node* last = NULL;
        Node* head = NULL;
        
        while(p){
            
            if(p->left){
                auto & cur = p->left;
                if(last)
                    last->next = cur;
                else{
                    // 没有前面的说明当前就是头
                    head =cur;
                }
                printf("%d ",cur->val);
                last = cur;
            }
            if(p->right){
                auto & cur = p->right;
                if(last)
                    last->next = cur;
                else{
                    // 没有前面的说明当前就是头
                    head =cur;
                }
                printf("%d ",cur->val);
                last = cur;
            }
            if(p->next){
                p=p->next;
            }else{
                // 换层
                p = head;
                head = NULL;
                last = NULL;
                printf("\n");
            }
            
        }
        return root;
        
    }
};
