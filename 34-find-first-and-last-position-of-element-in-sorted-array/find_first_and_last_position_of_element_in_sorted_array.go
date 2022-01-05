package main

import "fmt"

func searchRange(nums []int, target int) []int {
	// 首先是常规的二分查找，r在边界外，l=m+1 r=m
	l := 0
	nlen := len(nums)
	r := nlen
	m := nlen / 2
	ret := []int{-1, -1}
	for l < r {
		m = (l + r) / 2
		if target < nums[m] {
			r = m
		} else if target > nums[m] {
			l = m + 1
		} else {
			// 找到元素的话就往外找左右边界
			l = m
			// 这个一定会到左边界的外面，最小是-1
			for l > -1 && nums[l] == target {
				l--
			}
			r = m
			// 右边界的下一个，最大是nums的长度
			for r < nlen && nums[r] == target {
				r++
			}
			// 因为是边界外所以要l+1 和r-1，得到闭区间
			ret[0] = l + 1
			ret[1] = r - 1
			break
		}
	}
	return ret
}
func main() {
	testcases := [][]int{
		{5, 7, 7, 8, 8, 10},
		{5, 7, 7, 8, 8, 10},
		{5, 7, 7, 8, 8, 10},
		{5, 7, 7, 8, 8, 10},
		{5, 7, 7, 8, 8, 10},
		{},
	}
	targets := []int{8, 6, 5, 10, 100, 0}

	for i, nums := range testcases {
		t := targets[i]
		r := searchRange(nums, t)
		fmt.Println(r)

	}
}
