package construct_string_from_binary_tree

import (
	"fmt"
	"testing"
)

func TestTree2str(t *testing.T) {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(tree2str(p))

	p = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(tree2str(p))
}
