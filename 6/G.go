package main

import (
	"fmt"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func binsearch(from, to int64, f func(int64) bool) int64 {
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
	var w, h, n int64
	fmt.Scan(&w, &h, &n)

	res := binsearch(0, min(w, h)/2, func(d int64) bool {
		s1 := 2 * d
		return s1*(w+h-s1) > n
	})

	fmt.Println(res - 1)
}
