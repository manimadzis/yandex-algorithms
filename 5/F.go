package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	first  int
	second int
}

type Pairs struct {
	pairs []Pair
}

func (p Pairs) Less(i, j int) bool {
	p1 := p.pairs[i]
	p2 := p.pairs[j]

	return p1.first < p2.first || (p1.first == p2.first && p1.second < p2.second)
}

func (p Pairs) Len() int {
	return len(p.pairs)
}

func (p Pairs) Swap(i, j int) {
	p.pairs[i], p.pairs[j] = p.pairs[j], p.pairs[i]
}

func main() {

	var n, m int

	fmt.Scan(&n)

	powers := make(sort.IntSlice, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&powers[i])
	}

	fmt.Scan(&m)

	models := Pairs{
		pairs: make([]Pair, m),
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&models.pairs[i].second, &models.pairs[i].first)
	}

	sort.Sort(&models)
	powers.Sort()

	price := 0
	i := 0
	for _, power := range powers {
		for models.pairs[i].second < power {
			i++
		}

		price += models.pairs[i].first
	}

	fmt.Println(price)
}
