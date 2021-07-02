/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */



class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        // 用最小堆例子实现，给所有的链表建一个最小堆
        if(lists.size()==0){return  nullptr; }
        // 堆的排序算法，注意前面的空指针判断，要让空指针(空列表)在前面
        auto less= []( ListNode* &a, ListNode*  &b){ return (!a || !b)?(!a):(a->val > b->val);};

        make_heap(lists.begin(),lists.end(),less);
        ListNode* head=new ListNode();
        ListNode* ret=head;
        while(lists.size()){
            auto e = lists.front();
            pop_heap(lists.begin(),lists.end(),less);
            lists.pop_back();
            if(e==nullptr){continue;};
            ret->next = new ListNode(e->val);
            ret=ret->next;
            e=e->next;
            if(e){
                lists.push_back(e);
                push_heap(lists.begin(),lists.end(),less);
            }
        }
        return head->next;

    }   
};
