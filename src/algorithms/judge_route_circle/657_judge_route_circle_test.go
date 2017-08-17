package judge_route_circle

import (
	"fmt"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	moves := "UD"
	fmt.Println(judgeCircle(moves))

	moves = "LL"
	fmt.Println(judgeCircle(moves))
}
