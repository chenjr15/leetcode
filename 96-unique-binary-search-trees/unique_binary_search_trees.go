package leetcode96

// 递归
func numTrees(n int) int {
	if n < 1 {
		return 1
	}
	if n < 3 {
		return n
	}
	total := 0
	for i := 1; i <= n; i++ {

		total += numTrees(i-1) * numTrees(n-i)
	}
	return total

}
