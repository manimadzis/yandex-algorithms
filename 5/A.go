package main

import (
	"bufio"
	"fmt"
	"os"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n, k int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n)

	m := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &m[i])
	}

	fmt.Fscan(reader, &k)

	p := make([]int, k)

	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &p[i])
	}

	i, j := 0, 0
	diff := 10000001
	pair := [2]int{}

	for i < n && j < k && diff > 0 {
		delta := Abs(m[i] - p[j])
		if delta < diff {
			diff = delta
			pair[0] = m[i]
			pair[1] = p[j]
		}

		if m[i] >= p[j] {
			j++
		} else {
			i++
		}
	}

	fmt.Printf("%d %d\n", pair[0], pair[1])
}
