/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
 #include<climits>
class Solution {
public:
    long long last=LLONG_MIN;
    bool isValidBST(TreeNode* root) {
        // 用栈的非递归解法
        if (!root)return true;
        stack<TreeNode*> s;
        TreeNode * p =root;
        while(p||!s.empty()){
            // 元素非空
            if (p){
                // 入栈
                s.push(p);
                // 并往左走
                p = p->left;
                continue;
            }
            // 元素为空，但是栈非空
            if (!p && !s.empty()){
                // 出栈
                p=s.top();
                s.pop();
                // printf("%d,%d,\n",last,p->val);
                // 访问(比较值)
                if (last>=p->val){
                    return false;
                }
                last = p->val;
                // 往右走
                p = p->right;
            }
        }
        return true;
    }
};
