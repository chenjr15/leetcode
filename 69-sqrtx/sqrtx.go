package main

import "fmt"

func mySqrt(x int) int {
	l := 0
	r := x
	m := r / 2
	m2diff := m
	// 找左边界
	for l <= r {

		// 防止溢出
		m = l + (r-l)/2
		fmt.Println(l, r, m)
		m2diff = m*m - x
		if m2diff > 0 {
			r = m - 1
		} else if m2diff < 0 {
			l = m + 1
		} else {
			r = m
			break
		}
	}
	fmt.Println(l, r, m)
	return r
}

func main() {
	nums := []int{4, 8, 10, 5, 6, 7, 9}
	for _, x := range nums {

		r := mySqrt(x)
		fmt.Println("ret", x, r)
	}
}
