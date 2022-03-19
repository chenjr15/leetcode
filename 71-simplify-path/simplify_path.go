package leetcode71

import "strings"

func simplifyPath(path string) string {

	parts := strings.Split(path, "/")
	path_stack := make([]string, 0, len(parts))
	for _, elem := range parts {
		if elem == "" || elem == "." {
			continue
		} else if elem == ".." {
			if len(path_stack) > 0 {
				path_stack = path_stack[0 : len(path_stack)-1]
			}
		} else {
			path_stack = append(path_stack, elem)
		}
	}

	var b strings.Builder
	for _, elem := range path_stack {
		b.WriteRune('/')
		b.WriteString(elem)
	}
	if len(path_stack) == 0 {
		b.WriteRune('/')
	}

	return b.String()

}
