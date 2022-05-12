package main

import "fmt"

func median(a, b []int) int {
	i, j, c := 0, 0, 0
	L := len(a)

	for i < L && j < L && c < L-1 {
		if a[i] > b[j] {
			j++
		} else {
			i++
		}
		c++
	}

	for i < L && c < L-1 {
		i++
		c++
	}

	for j < L && c < L-1 {
		j++
		c++
	}

	if a[i] > b[j] {
		return b[j]
	} else {
		return a[i]
	}
}

func main() {

	var N, L int
	fmt.Scan(&N, &L)

	a := make([][]int, N)
	for i, _ := range a {
		a[i] = make([]int, L)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < L; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			fmt.Println(median(a[i], a[j]))
		}
	}
}
