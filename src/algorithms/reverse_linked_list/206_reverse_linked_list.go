package reverse_linked_list

// link https://leetcode.com/problems/reverse-linked-list/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	var new_head *ListNode

	for {
		new_node := &ListNode{
			Val:  head.Val,
			Next: new_head,
		}

		new_head = new_node

		if nil == head.Next {
			break
		}
		head = head.Next
	}

	return new_head
}
