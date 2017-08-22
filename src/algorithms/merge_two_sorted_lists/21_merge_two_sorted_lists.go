package merge_two_sorted_lists

// https://leetcode.com/problems/merge-two-sorted-lists/description/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}
	var head *ListNode
	var p *ListNode
	var curr_val int
	var is_head_init = false

	for {
		if l1.Val < l2.Val {
			curr_val = l1.Val

			l1 = l1.Next
		} else {
			curr_val = l2.Val

			l2 = l2.Next
		}

		node := &ListNode{
			Val: curr_val,
		}

		if is_head_init {
			p.Next = node
			p = p.Next
		} else {
			is_head_init = true

			head = node
			p = head
		}

		if nil == l1 {
			p.Next = l2
			break
		}
		if nil == l2 {
			p.Next = l1
			break
		}
	}

	return head
}
