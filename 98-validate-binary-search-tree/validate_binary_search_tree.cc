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
    // 改成int 指针会快挺多的，数据类型的优化不可忽视
    long  last =LONG_MIN; 
    bool isValidBST(TreeNode* root) {
        // 递归中序遍历，结果应是严格升序，即为bst
        // 将上一个值存在last成员变量中即可判断是否为bst
        if(!root){
            return true;
        }
        if (root->left && !isValidBST(root->left)){
            return false;
        }
        if ( root->val <= last){
            return false;
        }
        last=root->val;
        if (root->right && !isValidBST(root->right)){
            return false;
        }
        return true;
    }
};
