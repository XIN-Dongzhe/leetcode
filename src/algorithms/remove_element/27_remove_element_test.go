package remove_element

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{1, 2, 3, 3, 5, 7, 7, 7, 11, 15}
	val := 7
	fmt.Println(removeElement(nums, val), nums)

	nums = []int{3, 2, 2, 3}
	val = 3
	fmt.Println(removeElement(nums, val), nums)

	nums = []int{3, 2, 2, 3}
	val = 2
	fmt.Println(removeElement(nums, val), nums)

	nums = []int{3, 2, 2, 3}
	val = 4
	fmt.Println(removeElement(nums, val), nums)

	nums = []int{1}
	val = 1
	fmt.Println(removeElement(nums, val), nums)

	nums = []int{}
	val = 1
	fmt.Println(removeElement(nums, val), nums)
}
