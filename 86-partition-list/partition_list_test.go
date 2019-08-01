package leetcode86

import (
	"fmt"
	"testing"

	"github.com/chenjr15/golist/linkedlist"
)

func TestPartition(t *testing.T) {
	a := partition(linkedlist.MakeLinkedList([]int{1, 4, 3, 2, 5, 2}), 3)
	fmt.Println("124")
	cnt := 0
	for p := a; p != nil; p = p.Next {
		cnt++
		fmt.Println(p.Val)
		if cnt > 20 {
			break
		}
	}

}
