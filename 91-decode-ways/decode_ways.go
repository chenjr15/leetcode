package leetcode91

func numDecodings(s string) int {
	//
	n := len(s)
	if n == 0 {
		return 0
	}
	c := s[0]
	if c > 2 || c == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		c = s[i]

	}

	return dp[n-1]
}
