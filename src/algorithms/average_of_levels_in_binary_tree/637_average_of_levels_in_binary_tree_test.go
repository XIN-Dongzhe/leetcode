package average_of_levels_in_binary_tree

import (
	"fmt"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	root := &TreeNode{
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
	fmt.Println(averageOfLevels(root))

	root = &TreeNode{
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
	fmt.Println(averageOfLevels(root))
}
