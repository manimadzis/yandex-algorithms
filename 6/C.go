package main

import (
	"fmt"
)

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

	res := binsearch(0, 10000000000, func(size int64) bool {
		return (size/w)*(size/h) >= n
	})

	fmt.Println(res)
}
