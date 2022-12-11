package main

import (
	"fmt"
	"sort"
)

type monkey struct {
	items    []int
	op       func(int) int
	test     func(int) int
	nInspect int
}

var _monkeys []monkey = []monkey{
	{
		// Monkey 0
		items: []int{79, 98},
		op:    func(n int) int { return n * 19 },
		test: func(n int) int {
			if n%23 == 0 {
				return 2
			} else {
				return 3
			}
		},
	},
	{
		// Monkey 1
		items: []int{54, 65, 75, 74},
		op:    func(n int) int { return n + 6 },
		test: func(n int) int {
			if n%19 == 0 {
				return 2
			} else {
				return 0
			}
		},
	},
	{
		// Monkey 2
		items: []int{79, 60, 97},
		op:    func(n int) int { return n * n },
		test: func(n int) int {
			if n%13 == 0 {
				return 1
			} else {
				return 3
			}
		},
	},
	{
		// Monkey 3
		items: []int{74},
		op:    func(n int) int { return n + 3 },
		test: func(n int) int {
			if n%17 == 0 {
				return 0
			} else {
				return 1
			}
		},
	},
}

var monkeys []monkey = []monkey{
	{
		// Monkey 0
		items: []int{74, 64, 74, 63, 53},
		op:    func(n int) int { return n * 7 },
		test: func(n int) int {
			if n%5 == 0 {
				return 1
			} else {
				return 6
			}
		},
	},
	{
		// Monkey 1
		items: []int{69, 99, 95, 62},
		op:    func(n int) int { return n * n },
		test: func(n int) int {
			if n%17 == 0 {
				return 2
			} else {
				return 5
			}
		},
	},
	{
		// Monkey 2
		items: []int{59, 81},
		op:    func(n int) int { return n + 8 },
		test: func(n int) int {
			if n%7 == 0 {
				return 4
			} else {
				return 3
			}
		},
	},
	{
		// Monkey 3
		items: []int{50, 67, 63, 57, 63, 83, 97},
		op:    func(n int) int { return n + 4 },
		test: func(n int) int {
			if n%13 == 0 {
				return 0
			} else {
				return 7
			}
		},
	},
	{
		// Monkey 4
		items: []int{61, 94, 85, 52, 81, 90, 94, 70},
		op:    func(n int) int { return n + 3 },
		test: func(n int) int {
			if n%19 == 0 {
				return 7
			} else {
				return 3
			}
		},
	},
	{
		// Monkey 5
		items: []int{69},
		op:    func(n int) int { return n + 5 },
		test: func(n int) int {
			if n%3 == 0 {
				return 4
			} else {
				return 2
			}
		},
	},
	{
		// Monkey 6
		items: []int{54, 55, 58},
		op:    func(n int) int { return n + 7 },
		test: func(n int) int {
			if n%11 == 0 {
				return 1
			} else {
				return 5
			}
		},
	},
	{
		// Monkey 7
		items: []int{79, 51, 83, 88, 93, 76},
		op:    func(n int) int { return n * 3 },
		test: func(n int) int {
			if n%2 == 0 {
				return 0
			} else {
				return 6
			}
		},
	},
}

func solvePart1() {
	nRound := 20
	for round := 0; round < nRound; round++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				n := m.op(item)
				n /= 3
				dstMonkey := m.test(n)
				monkeys[dstMonkey].items = append(monkeys[dstMonkey].items, n)
				monkeys[i].nInspect++
			}

			monkeys[i].items = []int{}
		}
	}

	nInspects := []int{}
	for _, m := range monkeys {
		nInspects = append(nInspects, m.nInspect)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nInspects)))
	fmt.Println(nInspects[0] * nInspects[1])
}

func main() {
	solvePart1()
}
