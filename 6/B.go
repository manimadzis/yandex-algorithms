package main

import (
	"bufio"
	"fmt"
	"os"
)

func binsearch(arr []int, elem int) int {
	l, r := -1, len(arr)

	for r-l > 1 {
		mid := (r + l) / 2
		if arr[mid] >= elem {
			r = mid
		} else {
			l = mid
		}
	}

	return r
}

func main() {

	var n, k int

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	b := make([]int, k)

	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &b[i])
	}

	for i := 0; i < k; i++ {
		res := binsearch(a, b[i])

		if res < n && a[res] == b[i] {
			fmt.Println(a[res])
		} else {
			if res == 0 {
				fmt.Println(a[res])
			} else if res >= n {
				fmt.Println(a[res-1])
			} else {
				if a[res]-b[i] >= b[i]-a[res-1] {
					fmt.Println(a[res-1])
				} else {
					fmt.Println(a[res])
				}
			}
		}
	}
}
