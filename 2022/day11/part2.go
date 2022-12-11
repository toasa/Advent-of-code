package main

import (
	"fmt"
	"sort"
)

type monkey struct {
	items       []int
	op          func(int) int
	testDivisor int
	testThen    int
	testElse    int
	nInspect    int
}

var _monkeys []monkey = []monkey{
	{
		// Monkey 0
		items:       []int{79, 98},
		op:          func(n int) int { return n * 19 },
		testDivisor: 23,
		testThen:    2,
		testElse:    3,
	},
	{
		// Monkey 1
		items:       []int{54, 65, 75, 74},
		op:          func(n int) int { return n + 6 },
		testDivisor: 19,
		testThen:    2,
		testElse:    0,
	},
	{
		// Monkey 2
		items:       []int{79, 60, 97},
		op:          func(n int) int { return n * n },
		testDivisor: 13,
		testThen:    1,
		testElse:    3,
	},
	{
		// Monkey 3
		items:       []int{74},
		op:          func(n int) int { return n + 3 },
		testDivisor: 17,
		testThen:    0,
		testElse:    1,
	},
}

var monkeys []monkey = []monkey{
	{
		// Monkey 0
		items:       []int{74, 64, 74, 63, 53},
		op:          func(n int) int { return n * 7 },
		testDivisor: 5,
		testThen:    1,
		testElse:    6,
	},
	{
		// Monkey 1
		items:       []int{69, 99, 95, 62},
		op:          func(n int) int { return n * n },
		testDivisor: 17,
		testThen:    2,
		testElse:    5,
	},
	{
		// Monkey 2
		items:       []int{59, 81},
		op:          func(n int) int { return n + 8 },
		testDivisor: 7,
		testThen:    4,
		testElse:    3,
	},
	{
		// Monkey 3
		items:       []int{50, 67, 63, 57, 63, 83, 97},
		op:          func(n int) int { return n + 4 },
		testDivisor: 13,
		testThen:    0,
		testElse:    7,
	},
	{
		// Monkey 4
		items:       []int{61, 94, 85, 52, 81, 90, 94, 70},
		op:          func(n int) int { return n + 3 },
		testDivisor: 19,
		testThen:    7,
		testElse:    3,
	},
	{
		// Monkey 5
		items:       []int{69},
		op:          func(n int) int { return n + 5 },
		testDivisor: 3,
		testThen:    4,
		testElse:    2,
	},
	{
		// Monkey 6
		items:       []int{54, 55, 58},
		op:          func(n int) int { return n + 7 },
		testDivisor: 11,
		testThen:    1,
		testElse:    5,
	},
	{
		// Monkey 7
		items:       []int{79, 51, 83, 88, 93, 76},
		op:          func(n int) int { return n * 3 },
		testDivisor: 2,
		testThen:    0,
		testElse:    6,
	},
}

func solvePart1() {
	// Calc common divisor
	mod := 1
	for _, m := range monkeys {
		mod *= m.testDivisor
	}

	nRound := 10000
	for round := 0; round < nRound; round++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				n := m.op(item)
				var dstMonkey int
				if n%m.testDivisor == 0 {
					dstMonkey = m.testThen
				} else {
					dstMonkey = m.testElse
				}
				n %= mod
				monkeys[dstMonkey].items = append(monkeys[dstMonkey].items, n)
				monkeys[i].nInspect++
			}

			monkeys[i].items = []int{}
		}
	}

	nInspects := []int{}
	for i, m := range monkeys {
		fmt.Println("Monkey", i, "inspected", m.nInspect, "times")
		nInspects = append(nInspects, m.nInspect)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nInspects)))
	fmt.Println(nInspects[0] * nInspects[1])
}

func main() {
	solvePart1()
}
