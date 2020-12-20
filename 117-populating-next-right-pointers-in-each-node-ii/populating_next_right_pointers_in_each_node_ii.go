package leetcode117

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	n := root
	var head *Node = nil
	var last *Node = nil
	var cur *Node
	// 巧妙利用题设，借助 next指针进行额外空间复杂度为常数的层次遍历
	// (其实是将空间转移到了每个节点上，对于普通的二叉树是无法做到的),
	// 跟着next走就相当于实现了层次遍历，借助当前的行已有的next指针进行下一行的层次遍历
	// 最开始根节点的时候直接可以进行左右节点遍历，即为层次便利
	for n != nil {
		if n.Left != nil {
			cur = n.Left
			if last == nil {
				// 第一次访问
				head = cur
			} else {
				last.Next = cur
			}
			last = cur
		}
		if n.Right != nil {
			cur = n.Right
			if last == nil {
				// 第一次访问
				head = cur
			} else {
				last.Next = cur
			}
			last = cur
		}
		n = n.Next
		if n == nil {
			// 换行
			n = head
			head = nil
			last = nil
		}

	}

	return root

}
