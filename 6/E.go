package main

import (
	"fmt"
)

func binsearch(from, to int, f func(int) bool) int {
	l, r := from-1, to+1

	for r-l > 1 {
		mid := (r + l) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	sum := 2*a + 3*b + 4*c
	count := a + b + c

	res := binsearch(0, count, func(d int) bool {
		return 2*(sum+5*d) >= 7*(count+d)
	})

	fmt.Println(res)
}
