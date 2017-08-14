package remove_duplicates_from_sorted_array

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	cursor := 0
	dup_count := 0

	for i := 1; i < len(nums); i++ {
		if nums[cursor] == nums[i] {
			dup_count++
		} else {
			cursor++
			nums[cursor] = nums[i]
		}
	}

	return len(nums) - dup_count
}
