package leetcode344

func reverseString(s []byte) {
	// 前后指针在位交换
	var l, r int
	for r = len(s) - 1; l < r; {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
