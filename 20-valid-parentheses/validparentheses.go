package vaildparentheses

func isValid(s string) bool {
	stack := make([]rune, 0, len(s)/2)
	parentsMap := map[rune]rune{'}': '{', ')': '(', ']': '['}

	for _, c := range s {

		switch c {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			stack = append(stack, c)

		case ')':
			fallthrough
		case ']':
			fallthrough
		case '}':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != parentsMap[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0

}
