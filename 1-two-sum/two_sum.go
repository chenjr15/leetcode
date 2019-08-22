package leetcode1

func twoSum(nums []int, target int) []int {

	// 放map 里直接查找即可，这里不能有重复元素
	intmap := make(map[int]int)
	for i, n := range nums {
		intmap[n] = i
	}

	for i := 0; i < len(nums); i++ {

		if j, ok := intmap[target-nums[i]]; ok && j != i {
			return []int{i, j}
		}

	}

	return []int{}
}
