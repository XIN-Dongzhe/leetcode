package merge_two_binary_trees

// https://leetcode.com/problems/merge-two-binary-trees/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if nil == t1 && nil == t2 {
		return nil
	}
	if nil == t1 {
		return t2
	}
	if nil == t2 {
		return t1
	}

	return &TreeNode{
		Val: t1.Val + t2.Val,
		Left: mergeTrees(t1.Left, t2.Left),
		Right: mergeTrees(t1.Right, t2.Right),
	}
}
