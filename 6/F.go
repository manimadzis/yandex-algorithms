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
	var n, x, y int
	fmt.Scan(&n, &x, &y)

	var min_t, max_t int
	if x < y {
		min_t = x
		max_t = n * y
	} else {
		min_t = y
		max_t = n * x
	}

	res := binsearch(min_t, max_t, func(t int) bool {
		return 1+(t-min_t)/x+(t-min_t)/y >= n
	})

	fmt.Println(res)
}
