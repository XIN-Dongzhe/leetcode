package judge_route_circle

// link https://leetcode.com/problems/subtree-of-another-tree/description/
func judgeCircle(moves string) bool {
	horizontal := 0
	vertical := 0
	for _, move := range moves {
		switch move {
		case 'U':
			vertical += 1
			break
		case 'D':
			vertical -= 1
			break
		case 'L':
			horizontal += 1
			break
		case 'R':
			horizontal -= 1
			break
		}
	}

	return 0 == horizontal && 0 == vertical
}