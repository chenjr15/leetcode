package leetcode1614

func maxDepth(s string) int {

	deepth := 0
	maxDeepth := deepth
	for _, c := range s {
		if c == '(' {
			deepth++
			if deepth > maxDeepth {
				maxDeepth = deepth
			}
		} else if c == ')' {
			deepth--
		}
	}
	return maxDeepth

}
