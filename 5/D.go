package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, r int

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n, &r)

	m := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &m[i])
	}

	left, right := 0, 1
	count := 0

	for right < n {
		diff := m[right] - m[left]

		if diff > r {
			left++
			count += n - right
		} else {
			right++
		}
	}

	fmt.Println(count)
}
