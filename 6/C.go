package main

import (
	"fmt"
)

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func binsearch(from int64, to int64, function func(int64) bool) int64 {
	l, r := from-1, to+1

	for r-l > 1 {
		mid := (l + r) / 2
		if function(mid) {
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

	res := binsearch(1, max(w, h)*n, func(size int64) bool {
		return (size/w)*(size/h) >= n
	})

	fmt.Println(res)
}
