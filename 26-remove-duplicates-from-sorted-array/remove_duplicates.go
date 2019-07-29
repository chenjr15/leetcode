package leetcode26

// remove the duplicates in sorted array
// Tow Pointers , low for unique, high for all item
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)

	}
	low := 0
	high := 1
	for high < len(nums) {
		if nums[low] != nums[high] {
			low++
			nums[low] = nums[high]

		}
		high++

	}
	return low + 1
}
