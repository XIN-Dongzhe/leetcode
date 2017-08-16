package subtree_of_another_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// link https://leetcode.com/problems/subtree-of-another-tree/description/
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isIdentical(s, t) {
		return true
	} else {
		if nil == s.Left {
			if nil == s.Right {
				return false
			} else {
				return isSubtree(s.Right, t)
			}
		} else {
			if isSubtree(s.Left, t) {
				return true
			} else {
				return nil != s.Right && isSubtree(s.Right, t)
			}
		}
	}

}

func isIdentical(s *TreeNode, t *TreeNode) bool {
	if nil == s {
		if nil == t {
			return true
		} else {
			return false
		}
	} else {
		if nil == t {
			return false
		} else {
			return s.Val == t.Val && isIdentical(s.Left, t.Left) && isIdentical(s.Right, t.Right)
		}
	}
}