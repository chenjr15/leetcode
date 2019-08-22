package leetcode15

import "sort"

// 有重复元素没法用哈希表解
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	resultSet := make([][]int, 0)
	// result := make([]int,3)
	for i := 0; i < len(nums); i++ {
		results := twoSum(nums[i+1:], 0-nums[i])
		for _, result := range results {
			// result[0]=nums[i]
			sortTriple(result)
			if !isUnique(resultSet, result) {
				continue
			}
			resultSet = append(resultSet, result)
			// result = make([]int,3)
		}

	}
	return resultSet
}

// 排序三元组
func sortTriple(nums []int) {
	if nums[0] > nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
	}
	if nums[1] > nums[2] {
		nums[1], nums[2] = nums[2], nums[1]
	}
	if nums[0] > nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
	}
}

// 去重
func isUnique(resultSet [][]int, result []int) bool {
	for i := 0; i < len(resultSet); i++ {
		for ; i < len(resultSet) && resultSet[i][0] == result[0]; i++ {
			for ; i < len(resultSet) && resultSet[i][1] == result[1]; i++ {
				if resultSet[i][2] == result[2] {
					return false
				}
			}
		}
	}
	return true
}

// 要求元素升序, 可能会有多个结果!
func twoSum(nums []int, target int) (results [][]int) {
	var i, j, n int
	j = len(nums) - 1
	for i < j {
		n = nums[i] + nums[j]
		switch {
		case n > target:
			j--
		case n < target:
			i++
		case n == target:
			result := make([]int, 3)
			result[0] = -target
			result[1], result[2] = nums[i], nums[j]
			results = append(results, result)
			j--
		}
	}
	return results
}
