package leetcode725

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	var length int
	// 答案数组
	ret := make([]*ListNode, k)
	// 第一次遍历计算长度
	p := head
	for ; p != nil; p = p.Next {
		length++
	}
	if length == 0 {
		return ret
	}
	// 余数
	left := length % k
	// 基本长度
	baseLen := length / k
	// 每段链表的长度= 基本长度 + 余数>0?1:0, 余数用一次减一
	// p 用来遍历链表
	p = head
	// 用来断开链表
	last := head
	for partI := 0; partI < k; partI++ {
		// 把头指针赋值上
		ret[partI] = p
		// 加上基本长度
		for i := 0; i < baseLen; i++ {
			last = p
			p = p.Next
		}
		// 加上余数
		if left > 0 {
			last = p
			p = p.Next
			left--
		}
		// 断开链表
		last.Next = nil

	}
	return ret
}
