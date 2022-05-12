package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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
		POINT = 2
		END   = 3
	)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscan(reader, &n, &m)

	events := make(PairSlice, n*2+m)
	points := make([]int, m)
	line_count := make(map[int]int)

	for i := 0; i < n; i++ {
		var b, e int
		fmt.Fscan(reader, &b, &e)

		events[2*i] = [2]int{min(b, e), START}
		events[2*i+1] = [2]int{max(e, b), END}
	}

	for i := 2 * n; i < 2*n+m; i++ {
		var coord int

		fmt.Fscan(reader, &coord)

		events[i] = [2]int{coord, POINT}
		points[i-2*n] = coord
	}

	sort.Sort(events)

	lines := 0

	for _, event := range events {
		if event[1] == START {
			lines++
		}

		if event[1] == POINT {
			line_count[event[0]] = lines
		}

		if event[1] == END {
			lines--
		}
	}

	for _, point := range points {
		fmt.Fprintf(writer, "%d ", line_count[point])
	}
	writer.Flush()

}
