package leetcode673

func findNumberOfLIS(nums []int) int {

	max_len := 1
	max_cnt := 1

	// dp数组的含义是以下标i元素结尾的最长递增子序列的长度为dp[i]
	dp := make([]int, len(nums))
	// 那么我们知道dp[0]肯定是1，因为只有一个元素
	// 那么dp[1] 就是看nums[1]是不是大于nums[0]，如果大于nums[0]则说明nums[1]可以和nums[0]拼起来成为递增子序列。
	// 故nums[1]=2,否则为1
	dp[0] = 1
	// 加一个dup数组用来存个数
	dup := make([]int, len(nums))
	dup[0] = 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		dup[i] = 1
		// 计算nums[i] 是否可以和所有下标小于i的元素拼接成递增子序列。
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// 如果位置j的元素可以和位置i的元素拼成递增子序列则拼成的递增子序列长度为位置j的长度+1 即 dp[j]+1
				// 又因为我们定义的是最长的递增子序列长度，我们遍历所有小于该元素的元素求出最大的，
				// 因此每次都要和已有值比较，如果当前序列和前面的比比较大才替换进去。
				if dp[i] == (dp[j] + 1) {
					// 相等则添加重复个数即可
					dup[i] += dup[j]
				} else if dp[i] < (dp[j] + 1) {
					// 更大则替换掉重复个数
					dup[i] = dup[j]
					dp[i] = (dp[j] + 1)
				}
			}
		}

		// dp[i]计算完毕，保存全局最大值
		if max_len == dp[i] {
			max_cnt += dup[i]
		} else if dp[i] > max_len {
			max_len = dp[i]
			max_cnt = dup[i]
		}
		// fmt.Println("i=",i," dp[i] = ",dp[i]," dup[i]=",dup[i]," max_len=",max_len," max_cnt=",max_cnt)

	}
	return max_cnt
}
