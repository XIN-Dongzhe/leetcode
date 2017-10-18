package remove_duplicates_from_sorted_list

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	p := head
	for {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}

		if nil == p.Next {
			break
		}
	}

	return head
}
