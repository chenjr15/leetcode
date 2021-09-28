package leetcode639

func numDecodings(s string) int {
	var MOD int = 10e8 + 7
	n := len(s)
	if n == 0 {
		return 0
	}
	cur := s[0]
	pre := cur
	if cur == '0' {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	if cur == '*' {
		dp[1] = 9
	}
	for i := 2; i <= n; i++ {
		cur = s[i-1]
		pre = s[i-2]
		// 单独的情况，即不拼接
		if cur == '*' {
			// 1-9
			dp[i] += 9 * dp[i-1]
		} else if cur != '0' {
			dp[i] += dp[i-1]
		}
		// 和上个字符合并的情况，
		if cur == '*' && pre == '*' {
			// 两个都为*则有十五种情况
			dp[i] += dp[i-2] * 15
		} else if cur == '*' {
			if pre == '1' {
				// 1[1-9]
				dp[i] += 9 * dp[i-2]
			} else if pre == '2' {
				// 2[1-6]
				dp[i] += 6 * dp[i-2]
			} else {
				//不能接
				// [0,3-6]*
			}
		} else if pre == '*' {
			// 前面的*只能取1、2
			if cur < '7' {
				// 12都可以取
				// {1,2}.
				dp[i] += 2 * dp[i-2]
			} else {
				// 只能去1
				// 1.
				dp[i] += dp[i-2]
			}
		} else {
			// 都不是*
			if pre == '1' || (pre == '2' && cur < '7') {
				dp[i] += dp[i-2]
			}
		}
		dp[i] %= MOD
	}
	return dp[n]

}
