package subtree_of_another_tree

import (
	"fmt"
	"testing"
)

func TestIsSubtree(t *testing.T) {
	s := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	tr := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(isSubtree(s, tr))

	s = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 0,
				},
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	fmt.Println(isSubtree(s, tr))

	s = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 0,
				},
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	tr = &TreeNode{
		Val: 3,
	}
	fmt.Println(isSubtree(s, tr))
}
