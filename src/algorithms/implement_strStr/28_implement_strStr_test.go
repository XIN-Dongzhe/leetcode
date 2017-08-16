package implement_strStr

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	haystack := "bonjour"
	needle := "jour"
	fmt.Println(strStr(haystack, needle))

	haystack = "bonjour"
	needle = "hello"
	fmt.Println(strStr(haystack, needle))
}
