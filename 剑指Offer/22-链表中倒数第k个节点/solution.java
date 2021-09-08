/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode getKthFromEnd(ListNode head, int k) {
        ListNode fast = head;
        ListNode slow = head;
        while(k>1){
            fast=fast.next;
            --k;
        }
        while(fast.next!=null){
            fast = fast.next;
            slow = slow.next;
        }
        return slow;
   }
}
