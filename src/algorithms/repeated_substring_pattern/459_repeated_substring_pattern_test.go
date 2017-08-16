package repeated_substring_pattern

import (
	"fmt"
	"testing"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	s := "bb"
	fmt.Println(repeatedSubstringPattern(s))

	s = "aba"
	fmt.Println(repeatedSubstringPattern(s))

	s = "abcabcabcabc"
	fmt.Println(repeatedSubstringPattern(s))
}
