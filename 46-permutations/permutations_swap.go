package leetcode46

func permute(nums []int) [][]int {
	Size := [...]int{
		0,
		1,
		2,
		3 * 2,
		4 * 3 * 2,
		5 * 4 * 3 * 2,
		6 * 5 * 4 * 3 * 2,
	}
	nlen := len(nums)
	result := make([][]int, 0, Size[nlen])

	var backtrack func(sorted int)
	backtrack = func(sorted int) {
		if sorted == nlen {
			// 输出
			result = append(result, append([]int{}, nums...))
			return
		}
		backtrack(sorted + 1)
		for i := sorted + 1; i < nlen; i++ {
			// 交换
			nums[sorted], nums[i] = nums[i], nums[sorted]
			backtrack(sorted + 1)
			// 撤销交换
			nums[sorted], nums[i] = nums[i], nums[sorted]
		}

	}
	backtrack(0)
	return result

}
