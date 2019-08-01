package leetcode80

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	if len(nums) == 3 {
		if nums[0] == nums[1] && nums[1] == nums[2] {
			return 2
		} else {
			return 3
		}
	}
	slow := 2
	for fast := 2; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
