package leetcode509

func fib(n int) int {
	// 斐波那契
	result := []int{0, 1}
	if n < 2 {
		return result[n]
	}
	for i := 2; i <= n; i++ {
		current := result[0] + result[1]
		result[0] = result[1]
		result[1] = current
	}
	return result[1]
}
