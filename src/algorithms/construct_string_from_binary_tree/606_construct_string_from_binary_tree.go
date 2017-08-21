package construct_string_from_binary_tree

import (
	"strconv"
	"bytes"
)

// https://leetcode.com/problems/construct-string-from-binary-tree/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if nil == t {
		return ""
	}
	var s bytes.Buffer
	s.WriteString(strconv.Itoa(t.Val))

	if nil == t.Left && nil == t.Right {
		return s.String()
	}

	s.WriteString("(")
	s.WriteString(tree2str(t.Left))
	s.WriteString(")")

	if nil != t.Right {
		s.WriteString("(")
		s.WriteString(tree2str(t.Right))
		s.WriteString(")")
	}

	return s.String()
}
