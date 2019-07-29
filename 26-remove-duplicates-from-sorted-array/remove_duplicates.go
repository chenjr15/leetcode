package leetcode26

//remove the duplicates in sorted array
func removeDuplicates(nums []int) int {
	count := len(nums)
	if len(nums) == 1 {
		return 1

	}
	for i := 0; i < count-1; {
		if nums[i] == nums[i+1] {
			nextdiff := i + 2
			for nextdiff < count {
				if nums[i] != nums[nextdiff] {
					break
				}
				nextdiff++
			}
			copy(nums[i+1:], nums[nextdiff:])
			i = i + 1
			count = count - (nextdiff - i)

		} else {
			i++

		}

	}
	return count

}
