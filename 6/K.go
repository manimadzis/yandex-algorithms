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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func median(a, b []int) int {
	return binsearch(min(a[0], b[0]), max(a[len(a)-1], b[len(a)-1]), func(x int) bool {
		var l, r int

		l = binsearch(0, len(a)-1, func(i int) bool {
			return a[i] >= x
		})

		r = binsearch(0, len(a)-1, func(i int) bool {
			return a[i] > x
		})
		r -= 1

		if l > len(a)-1 {
			return false
		}
	})
}

func main() {

	var N, L int
	fmt.Scan(&N, &L)

	list := make([][]int, N)
	for i, _ := range list {
		list[i] = make([]int, L)
	}

	for i := 0; i < N; i++ {
		var x1, d1, a, c, m int
		fmt.Scan(&x1, &d1, &a, &c, &m)
		list[i][0] = x1
		for j := 1; j < L; j++ {
			list[i][j] = list[i][j-1] + d1
			d1 = (a*d1 + c) % m
		}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			fmt.Println(median(list[i], list[j]))
		}
	}
}
