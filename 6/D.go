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

func binsearch(from, to int64, function func(int64) bool) int64 {
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
	var n, a, b, w, h int64
	fmt.Scan(&n, &a, &b, &w, &h)

	max_d := max(max(w-a, w-b), max(h-a, h-b))

	res := binsearch(0, max_d, func(d int64) bool {
		new_n_1, new_n_2 := (w/(a+2*d))*(h/(b+2*d)), (h/(a+2*d))*(w/(b+2*d))
		return max(new_n_1, new_n_2) < n
	})

	fmt.Println(res - 1)
}
