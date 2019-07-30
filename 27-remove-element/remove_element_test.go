package leetcode27

import "testing"

func SlicePrefixEquals(sl1, sl2 []int) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i, _ := range sl1 {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}
func TestRemoveElement(t *testing.T) {
	testcases := []struct {
		data   []int
		val    int
		answer []int
	}{
		{
			[]int{0, 0, 1, 1, 2, 2, 2, 3, 3, 4}, 2, []int{0, 0, 1, 1, 3, 3, 4},
		},
		{
			[]int{0, 0, 0, 0, 0, 0, 1}, 0, []int{1},
		},
	}
	for i, testdata := range testcases {
		uni := removeElement(testdata.data, testdata.val)
		if !SlicePrefixEquals(testdata.answer, testdata.data[:uni]) {
			t.Errorf("%v/%v Failed, expected:%v, got: %v, len %v", i+1, len(testcases), testdata.answer, testdata.data[:uni], uni)
		} else {
			t.Logf("%v/%v Passed, %v ,%v", i+1, len(testcases), testdata.data[:uni], uni)
		}

	}
}
