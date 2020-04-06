package leetcode1137

func tribonacci(n int) int {
	// 三元斐波那契
	result := []int{0, 1, 1}
	if n < 3 {
		return result[n]
	}
	for i := 3; i <= n; i++ {
		current := result[0] + result[1] + result[2]
		result[0] = result[1]
		result[1] = result[2]
		result[2] = current
	}
	return result[2]
}
