package leetcode70

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	result := [2]int{1, 1}
	current := 0
	for i := 2; i <= n; i++ {
		current = result[0] + result[1]
		result[0] = result[1]
		result[1] = current
	}
	return current
}
