/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    struct ListNode* head=malloc(sizeof(struct ListNode));
    struct ListNode* current=head;
    struct ListNode* next;
    int remain = 0;
    while(l1||l2||remain){
        next =malloc(sizeof(struct ListNode));
        next->next = NULL;
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
    free(head);
    return next;
}
