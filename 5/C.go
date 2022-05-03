package main

import "fmt"

func height(trace []int, left, right int) int {
	return trace[right] - trace[left]
}

func main() {

	var n int
	fmt.Scan(&n)
	trace_forward := make([]int, n)
	trace_back := make([]int, n)

	var prev, cur, x int
	fmt.Scan(&x, &prev)

	for i := 1; i < n; i++ {
		fmt.Scan(&x, &cur)

		if prev < cur {
			trace_forward[i] = trace_forward[i-1] + cur - prev
			trace_back[i] = trace_back[i-1]
		} else {
			trace_forward[i] = trace_forward[i-1]
			trace_back[i] = trace_back[i-1] + prev - cur
		}

		prev = cur
	}

	var m, left, right int
	fmt.Scan(&m)
	desc := make([][2]int, m)

	for i := 0; i < m; i++ {
		fmt.Scan(&left, &right)
		desc[i][0] = left
		desc[i][1] = right
	}

	for i := 0; i < m; i++ {
		if desc[i][0] < desc[i][1] {
			fmt.Println(height(trace_forward, desc[i][0]-1, desc[i][1]-1))
		} else {
			fmt.Println(height(trace_back, desc[i][1]-1, desc[i][0]-1))

		}
	}

}
