package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n, &k)
	m := make([]int, n)
	count := make(map[int]int)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &m[i])
		count[m[i]]++
	}

	left := 0
	for {
		if count[m[left]] > 1 {
			count[m[left]]--
			left++
		} else {
			break
		}
	}

	right := n - 1
	for {
		if count[m[right]] > 1 {
			count[m[right]]--
			right--
		} else {
			break
		}
	}
	fmt.Println(left+1, " ", right+1)
}
