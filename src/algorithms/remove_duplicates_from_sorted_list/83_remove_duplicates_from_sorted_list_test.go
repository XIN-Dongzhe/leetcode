package remove_duplicates_from_sorted_list

import (
	"testing"
	"fmt"
)

func TestDeleteDuplicates(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	}

	new_list := deleteDuplicates(list)

	for {
		fmt.Println(new_list.Val)
		if nil == new_list.Next {
			break
		}

		new_list = new_list.Next
	}
}
