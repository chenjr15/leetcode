package leetcode167

func twoSum(numbers []int, target int) []int {

	length := len(numbers)
	l := 0
	r := length
	// 二分查找
	search := func(n int) (index int) {
		index = -1
		l = 0
		r = length
		for i := length / 2; r > l; i = (l + r) / 2 {
			switch {
			case numbers[i] == n:
				return i
			case numbers[i] > n:
				r = i
			case numbers[i] < n:
				l = i + 1
			}
		}
		return -1
	}
	// 遍历到大于等于当前的元素即可，因为有序所以更大的不可能存在可能的组合
	for i := 0; numbers[i] <= target; i++ {
		idx := search(target - numbers[i])
		if idx != -1 && idx != i {

			return []int{i + 1, idx + 1}
		}
	}
	return []int{-1, -1}

}
