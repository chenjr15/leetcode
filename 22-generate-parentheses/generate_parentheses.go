package main

import "fmt"

func main() {
	fmt.Println("n =3 ")
	for _, s := range generateParenthesis(3) {
		fmt.Println(s)
	}
}
func generateParenthesis(n int) []string {
	// `left` and `right` means the left and right parentheses that could be used
	// left must be placed before right
	var result = make([]string, 0)
	seq := make([]byte, n+n)
	var addParentheses func(left int, right int)
	addParentheses = func(left int, right int) {
		if right == 0 {
			// all closed , time to return
			result = append(result, string(seq))
		}
		// first place left
		if left != 0 {
			seq[n+n-left-right] = '('
			addParentheses(left-1, right)
		}
		// at the first time, right equals left, this branch will not enter yet.
		// after enter the branch above(which insert left), the right remains more than left,
		// we can insert right now
		if right > left {
			seq[n+n-left-right] = ')'
			addParentheses(left, right-1)
		}
	}
	addParentheses(n, n)
	return result

}
