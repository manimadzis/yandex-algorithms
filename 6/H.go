package main

import "fmt"

func binsearch(from, to int, f func(int) bool) int {
	l, r := from-1, to+1

	for r-l > 1 {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid
		}
	}

	return r
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	res := binsearch(1, 10_000_000, func(length int) bool {
		count := 0
		for _, item := range a {
			count += item / length
		}
		return count < k
	})

	fmt.Println(res - 1)
}
