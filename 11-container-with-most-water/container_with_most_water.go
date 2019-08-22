package leetcode11

func maxArea(height []int) (max int) {
	var (
		left  int
		right int
	)
	// 用左右指针分别表示容器两边
	right = len(height) - 1
	var current int
	for left < right {
		// 每次计算当前的容量, 并收缩较低的边
		if height[left] < height[right] {
			current = (height[left]) * (right - left)
			left++
		} else {
			current = (height[right]) * (right - left)
			right--
		}
		if max < current {
			max = current
		}
	}

	return max

}
