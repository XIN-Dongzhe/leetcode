package search_insert_position

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target), nums)

	nums = []int{1, 3, 5, 6}
	target = 2
	fmt.Println(searchInsert(nums, target), nums)

	nums = []int{1, 3, 5, 6}
	target = 7
	fmt.Println(searchInsert(nums, target), nums)

	nums = []int{1, 3, 5, 6}
	target = 0
	fmt.Println(searchInsert(nums, target), nums)
}
