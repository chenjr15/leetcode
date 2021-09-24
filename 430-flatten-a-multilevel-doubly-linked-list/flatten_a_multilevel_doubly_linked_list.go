package leetcode430

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	p := root
	tails := make([]*Node, 0)
	last := root
	for p != nil || len(tails) != 0 {
		if p == nil {
			// 走到子链表的结尾，把上级链表拼接回去
			// len(tails) !=0
			p = last
			tail := tails[len(tails)-1]
			tails = tails[0 : len(tails)-1]

			p.Next = tail
			p.Next.Prev = p

		}
		if p.Child != nil {
			if p.Next != nil {
				// 如果下一个节点不空则把下一个节点存起来
				tails = append(tails, p.Next)
			}
			// 把子节点接到下一个节点的位置
			p.Next = p.Child
			p.Child.Prev = p
			p.Child = nil
		}
		last = p
		p = p.Next
	}
	return root

}
