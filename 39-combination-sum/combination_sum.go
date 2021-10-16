package leetcode39

func combinationSum(candidates []int, target int) [][]int {
	curSum := 0
	clen := len(candidates)
	results := make([][]int, 0)
	combination := make([]int, 0)
	var temp []int = nil
	deepth := 1
	// 回溯法,其实就是深搜
	var backtrack func(lastIdx int)
	backtrack = func(lastIdx int) {
		if deepth >= 150 {
			// 题目给定长度最大为150,防止无解
			return
		}
		// 这里直接从上一个元素开始来加，保证索引小的不会添加到索引大的后面,lastIdx 是上一次的索引
		// 注意可以无限制重复添加某个数字这个条件，所以这里i=lastIdx，如果不能重复就是i=lastIdx+1
		for i := lastIdx; i < clen; i++ {
			curSum += candidates[i]
			// 提前剪枝
			if curSum <= target {
				combination = append(combination, candidates[i])
				deepth++
				// 题目给定候选数字至少为1, 因此这里可以直接剪枝
				if curSum == target {
					// 刚好等于targe,可以加入结果slice
					temp = make([]int, len(combination))
					copy(temp, combination)
					results = append(results, temp)

				} else {
					backtrack(i)
				}
				deepth--
				combination = combination[0 : len(combination)-1]
			}

			curSum -= candidates[i]
		}
	}
	backtrack(0)
	return results
}
