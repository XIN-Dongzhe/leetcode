package min_stack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := Constructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	min1 := minStack.GetMin()
	minStack.Pop()
	top := minStack.Top()
	min2 := minStack.GetMin()

	fmt.Println(min1)
	fmt.Println(top)
	fmt.Println(min2)

	minStack = Constructor()

	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)

	minStack.Top()
	minStack.Pop()
	minStack.GetMin()
	minStack.Pop()
	minStack.GetMin()
	minStack.Pop()

	minStack.Push(2147483647)
	minStack.Top()
	minStack.GetMin()

	minStack.Push(-2147483648)

	minStack.Top()
	minStack.GetMin()
	minStack.Pop()
	minStack.GetMin()
}
