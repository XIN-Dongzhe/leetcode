package sum_of_square_numbers

import (
	"fmt"
	"testing"
)

func TestJudgeSquareSum(t *testing.T) {
	c := 5
	fmt.Println(judgeSquareSum(c))

	c = 2
	fmt.Println(judgeSquareSum(c))

	c = 3
	fmt.Println(judgeSquareSum(c))

	c = 4
	fmt.Println(judgeSquareSum(c))

	c = 3823622865
	fmt.Println(judgeSquareSum(c))
}
