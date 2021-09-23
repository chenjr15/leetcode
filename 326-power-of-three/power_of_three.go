package leetcode326

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	// 3的幂一定是3的倍数，但是3的倍数不一定是3的幂，如果不是3的倍数，一定不是3的幂。
	// 3的幂 => 3的倍数
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	// 如果是3的幂最后一定会得到1
	return n == 1
}
