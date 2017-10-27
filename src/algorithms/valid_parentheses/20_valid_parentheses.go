package valid_parentheses

// https://leetcode.com/problems/valid-parentheses/description/
func isValid(s string) bool {
	str_len := len(s)
	if 1 == str_len % 2 {
		return false
	}

	char_stack := []rune{}
	for _, char := range s {
		if 40 == char || 123 == char || 91 == char {
			char_stack = append(char_stack, char)
		} else {
			if len(char_stack) < 1 {
				return false
			}
			left_par := char_stack[len(char_stack) - 1]
			if (40 == left_par && 41 == char) ||
				(123 == left_par && 125 == char) ||
				(91 == left_par && 93 == char){
				char_stack = char_stack[:len(char_stack) - 1]
			} else {
				return false
			}
		}
	}

	if len(char_stack) > 0 {
		return false
	} else {
		return true
	}
}
