package leetcode24

import "testing"
import "github.com/chenjr15/golist/linkedlist"

func TestSwapPairs(t *testing.T) {
	test := func(list *ListNode) {

		t.Logf("Before Swap %v", list)
		rlist := swapPairs(list)
		t.Logf("After Swap %v", rlist)

	}
	datas := [][]int{
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3},
		[]int{1, 2},
		[]int{1},
		[]int{},
	}
	for _, l := range datas {
		list := linkedlist.MakeLinkedList(l)

		test(list)
	}
}
