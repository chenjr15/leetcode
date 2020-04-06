package leetcode746

func minCostClimbingStairs(cost []int) int {
	// len(cost) >=2 , it's given condition
	result := make([]int, len(cost)+1)
	result[0] = cost[0]
	result[1] = cost[1]

	i := 2
	for ; i < len(cost); i++ {
		if result[i-1] < result[i-2] {
			result[i] = cost[i] + (result[i-1])

		} else {
			result[i] = cost[i] + (result[i-2])
		}

	}
	// last statir
	if result[i-1] < result[i-2] {
		result[i] = (result[i-1])

	} else {
		result[i] = (result[i-2])
	}
	return result[len(cost)]
}
