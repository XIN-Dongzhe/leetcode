package repeated_substring_pattern

import (
	"strings"
)

// link https://leetcode.com/problems/repeated-substring-pattern/description/
func repeatedSubstringPattern(s string) bool {
	str_len := len(s)

	for i := str_len/2 - 1; i >= 0; i-- {
		if 0 == str_len%(i+1) {
			cmp_str := strings.Repeat(s[:i+1], str_len/(i+1))
			if 0 == strings.Compare(s, cmp_str) {
				return true
			}
		}
	}

	return false
}
