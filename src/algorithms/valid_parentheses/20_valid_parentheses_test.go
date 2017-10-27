package valid_parentheses

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	str := "()"
	fmt.Println(isValid(str))

	str = "()[]{}"
	fmt.Println(isValid(str))

	str = "(]"
	fmt.Println(isValid(str))

	str = "([)]f"
	fmt.Println(isValid(str))
}
