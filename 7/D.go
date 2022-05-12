package main

import (
	"fmt"
	"sort"
)

type Type int

const (
	INCOME Type = iota
	OUTCOME
)

type Event struct {
	Type Type
	Time int
}

type Events []Event

func (e Events) Less(i, j int) bool {
	return e[i].Time < e[j].Time || (e[i].Time == e[j].Time && e[i].Type < e[j].Type)
}

func (e Events) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Events) Len() int {
	return len(e)
}

func main() {

	var n int

	fmt.Scan(&n)
	events := make(Events, 2*n)
	true_len := 0

	for i := 0; i < n; i++ {
		var intime, outtime int
		fmt.Scan(&intime, &outtime)

		if outtime-intime >= 5 {
			events[true_len].Type = INCOME
			events[true_len].Time = intime
			true_len++
			events[true_len].Type = OUTCOME
			events[true_len].Time = outtime
			true_len++
		}
	}

	sort.Sort(events)
	events = events[:true_len]
	count := 0
	maxcount, prevmaxcount := 0, 0
	maxtime, prevmaxtime := 0, 0

	for _, event := range events {
		if event.Type == INCOME {
			count++
		} else if event.Type == OUTCOME {
			if count > maxcount {
				prevmaxcount = maxcount
				maxcount = count
				prevmaxtime = maxtime
				maxtime = event.Time
			} else if count > prevmaxcount {
				prevmaxcount = count
				prevmaxtime = event.Time
			}
		}
	}
}
