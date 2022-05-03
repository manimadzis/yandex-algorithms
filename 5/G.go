package main

import "fmt"

func main() {

	var n, k int

	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	i := 0
	j := 0
	same := 1
	mul := 1
	count := 0

	for i < n {

		for j < n-1 && a[j] <= a[i]*k {
			j++
			if a[j] == a[j-1] {
				same++
			} else {
				mul *= same
				same = 1
			}
		}

		fmt.Println(mul)
		n := j - i

		count += n * (n - 1) * (n - 2) / mul

		cur := a[i]
		for i < n && a[i] == cur {
			i++
		}
	}

	fmt.Println(count)
}
