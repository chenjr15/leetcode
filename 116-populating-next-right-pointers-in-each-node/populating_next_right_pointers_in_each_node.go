package leetcode116

import "container/list"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// LevelOrderTraversal 层次遍历
// 用队列实现，先把头节点入队， 然后取队头，访问队头，再把对头的左右孩子入队, 循环至队列空
// level 从1 开始
func LevelOrderTraversal(t *Node, visit func(*Node, int)) {
	if t == nil {
		return
	}
	level := 0
	lastLevel := 0
	queue := list.New()
	if t != nil {
		queue.PushBack(t)
		lastLevel = 1
	}

	for lastLevel != 0 {
		thisLevel := lastLevel
		lastLevel = 0
		for thisLevel != 0 {
			t = (queue.Remove(queue.Front())).(*Node)
			thisLevel--
			if t.Left != nil {
				queue.PushBack(t.Left)
				lastLevel++
			}
			if t.Right != nil {
				queue.PushBack(t.Right)
				lastLevel++
			}
			visit(t, level)
		}
		level++

	}

}
func connect(root *Node) *Node {
	// 套以前的层序遍历模板
	lastLevel := 0
	var lastNode *Node
	var visit = func(n *Node, level int) {
		if lastNode != nil {
			// fmt.Println(lastNode.Val,n.Val,level)
			// 同一层内将上一个节点的Next指向当前节点
			if lastLevel == level {
				lastNode.Next = n
			}
		}
		lastNode = n
		lastLevel = level
	}
	LevelOrderTraversal(root, visit)
	return root
}
