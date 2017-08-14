package implement_strStr

// https://leetcode.com/problems/implement-strstr/description/
func strStr(haystack string, needle string) int {
	needle_len := len(needle)
	cmp_len := len(haystack) - needle_len

OUTER_LOOP:
	for i := 0; i <= cmp_len; i++ {
		for j := 0; j < needle_len; j++ {
			if haystack[i+j] != needle[j] {
				continue OUTER_LOOP
			}
		}
		return i
	}

	return -1
}
