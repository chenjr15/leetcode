/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        
        ListNode* head=new ListNode(0);
        ListNode* current=head;
        ListNode* next;
        int remain = 0;
        while(l1||l2||remain){
            next =new ListNode(0);
            if(l1&&l2){
                // 两个链表都没有走完
                next->val = l1->val+l2->val+remain;
                l1=l1->next;
                l2=l2->next;
            }else if(l1){
                //剩下L2没有走完
                next->val = l1->val+remain;
                l1=l1->next;
            }else if(l2){
                //剩下L2没有走完
                 next->val = l2->val+remain;
                l2=l2->next;
            }else{
                //还有进位的数字
                next->val = remain;
            }
            remain = (next->val)/10;
            next->val = (next->val)%10;
            current->next = next;
            current = next;
        }
        next = head->next;
        delete head;
        return next;
    }
};
