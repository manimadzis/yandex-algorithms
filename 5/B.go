package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int

	fmt.Fscan(reader, &n, &k)

	prefixsum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		var new int
		fmt.Fscan(reader, &new)
		prefixsum[i] = prefixsum[i-1] + new
	}

	left, right := 0, 1
	count := 0
	for right <= n {
		diff := prefixsum[right] - prefixsum[left]
		if diff == k {
			count++
			left++
			right++
		} else if diff > k {
			left++
		} else {
			right++
		}
	}

	fmt.Println(count)
}
