package leetcode27

/* removeElement
 * Given an array nums and a value val,
 * remove all instances of that value in-place
 * and return the new length.
 * Solution: use two pointers, fast ponter to travel all elements.
 * Once found value different from specific val,
 * copy the value in fast to slow pointer position and increase slow.
 * Finally, return the value of slow pointer.
 */
func removeElement(nums []int, val int) int {
	slow := 0

	for fast := slow; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}

	}
	return slow

}
