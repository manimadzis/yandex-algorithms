package main

import "fmt"

func main() {

	var n, k int

	fmt.Scan(&n, &k)

	var s string
	fmt.Scan(&s)

	d := make(map[byte]int)
	max_index, max_len := 0, 0
	i, j := 0, 0

	for j < n {
		for j < n && d[s[j]] < k {
			d[s[j]]++
			j++
		}

		if j - i > max_len {
			max_len = j - i
			max_index = i
		}

		d[s[i]]--
		i++
	}

	fmt.Printf("%d %d\n", max_len, max_index + 1)
}	
