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
	var result = make([]string, 0)
	var addParentheses func(inProgress string, left int, right int)
	addParentheses = func(inProgress string, left int, right int) {
		if right == 0 {
			// all closed , time to return
			result = append(result, inProgress)
		}
		if left != 0 {
			addParentheses(inProgress+"(", left-1, right)
		}
		if right > left {
			addParentheses(inProgress+")", left, right-1)
		}
	}
	addParentheses("", n, n)
	return result

}
