package main

import (
	"fmt"
	"sort"
)

type PairSlice [][2]int
type Pair [2]int

func (p PairSlice) Len() int {
	return len(p)
}

func (p PairSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairSlice) Less(i, j int) bool {
	return p[i][0] < p[j][0] || (p[i][0] == p[j][0] && p[i][1] < p[j][1])
}

func main() {
	var n, m int
	const (
		START = 1
		END   = 2
	)
	fmt.Scan(&n, &m)

	events := make(PairSlice, m*2)
	for i := 0; i < m; i++ {
		var b, e int
		fmt.Scan(&b, &e)

		events[2*i] = [2]int{b, START}
		events[2*i+1] = [2]int{e, END}
	}

	sort.Sort(events)
	visions := 0
	nocheaters := 0
	prev := 0

	for _, event := range events {
		if event[1] == START {
			visions++
		}

		if visions > 1 {
			nocheaters += event[0] - prev
		}

		if visions == 1 && event[1] == END {
			nocheaters++
		}

		if event[1] == END {
			visions--
		}

		prev = event[0]
	}

	fmt.Println(n - nocheaters)
	fmt.Println(nocheaters)
}
