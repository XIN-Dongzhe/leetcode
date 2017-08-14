package search_insert_position

// https://leetcode.com/problems/search-insert-position/description/
func searchInsert(nums []int, target int) int {
	if nums[0] > target {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == target || (nums[i] > target && nums[i-1] < target) {
			return i
		}
	}

	return len(nums)
}
