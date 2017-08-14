package remove_element

// https://leetcode.com/problems/remove-element/description/
func removeElement(nums []int, val int) int {
	cursor := 0
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[cursor] = nums[i]
			cursor++
		}
	}

	return cursor
}
