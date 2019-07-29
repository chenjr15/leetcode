package leetcode26

import (
	"fmt"
	"testing"
)

func SlicePrefixEquals(sl1, sl2 []int) bool {
	if len(sl1) > len(sl2) {
		return false
	}
	for i, _ := range sl1 {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}

func TestRemoveDuplicates(t *testing.T) {
	testcases := []struct {
		data   []int
		answer []int
	}{
		{
			[]int{1}, []int{1},
		},
		{
			[]int{1, 1}, []int{1},
		},
		{
			[]int{1, 1, 1}, []int{1},
		},
		{
			[]int{0, 0, 1, 1, 2, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4},
		},
		{
			[]int{0, 0, 0, 0, 0, 0, 1}, []int{0, 1},
		},
	}
	for i, testdata := range testcases {
		fmt.Printf("\nStart test %v \n\n", i)
		uni := removeDuplicates(testdata.data)
		if uni != len(testdata.answer) && SlicePrefixEquals(testdata.answer, testdata.data) {
			t.Errorf("%v/%v Failed, expected:%v, got: %v, len %v", i+1, len(testcases), testdata.answer, testdata.data, uni)
		} else {
			t.Logf("%v/%v Passed, %v ,%v", i+1, len(testcases), testdata.answer, uni)
		}

	}

}
