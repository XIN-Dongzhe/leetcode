package same_tree

// https://leetcode.com/problems/same-tree/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if nil == p {
		if nil == q {
			return true
		} else {
			return false
		}
	} else {
		if nil == q {
			return false
		} else {
			return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}
}