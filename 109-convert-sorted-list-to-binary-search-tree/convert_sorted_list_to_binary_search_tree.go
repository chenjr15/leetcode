package leetcode109

import "github.com/chenjr15/golist/linkedlist"

type ListNode = linkedlist.ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	// 要平衡就得保证左右节点数量一样多 ---> 根节点是中位数
	// 因为链表是排序的，所以直接每次取中点即可找到中位数
	// 可以找到中点，然后递归建左右子树， 这样问题就变成链表找中点和递归建树了
	// 递归很简单，但问题也很明显，首先调用栈会比较深，然后链表找中点每次都得快慢指针遍历

}
