package two_sum

// link https://leetcode.com/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSum1(nums []int, target int) []int {
	nums_map := make(map[int]int)

	for i, value := range nums {
		if j, ok := nums_map[target-value]; ok {
			return []int{i, j}
		}
		nums_map[value] = i
	}

	return []int{}
}
