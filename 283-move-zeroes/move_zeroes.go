package leetcode283

// simple two pointers
func moveZeroes(nums []int) {
	f, s := 0, 0
	for ; f < len(nums); f++ {
		if nums[f] != 0 {
			nums[s] = nums[f]
			s++
		}
	}
	for ; s < len(nums); s++ {
		nums[s] = 0
	}
}
