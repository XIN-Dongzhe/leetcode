package remove_duplicates_from_sorted_array

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 2, 3, 3, 5, 7, 7, 7, 11, 15}
	fmt.Println(removeDuplicates(nums), nums)

	nums = []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums), nums)

	nums = []int{1}
	fmt.Println(removeDuplicates(nums), nums)

	nums = []int{}
	fmt.Println(removeDuplicates(nums), nums)
}
