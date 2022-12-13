package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInputFromFile(filename string) []string {
	lines := []string{}

	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	s := bufio.NewScanner(fp)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	if err = s.Err(); err != nil {
		panic(err)
	}

	return lines
}

type heightMap struct {
	start  [2]int
	goal   [2]int
	height [][]int
}

type move struct {
	pos    [2]int
	nsteps int
}

func createHeightMap(input []string) heightMap {
	hm := heightMap{}

	height := make([][]int, len(input))
	for i, line := range input {
		for j, c := range line {
			if c == 'S' {
				hm.start[0] = i
				hm.start[1] = j
				height[i] = append(height[i], 0)
			} else if c == 'E' {
				hm.goal[0] = i
				hm.goal[1] = j
				height[i] = append(height[i], 25)
			} else {
				height[i] = append(height[i], int(c-'a'))
			}
		}
	}
	hm.height = height

	return hm
}

var fewestSteps [][]int

const LIMIT = 500

func getValidNextMoves(m move, height [][]int) []move {
	nexts := []move{}

	canMove := func(cur, next [2]int) bool {
		curH := height[cur[0]][cur[1]]
		nextH := height[next[0]][next[1]]
		if curH >= nextH || curH+1 == nextH {
			return true
		}
		return false
	}

	// Check up
	if m.pos[0]-1 >= 0 && canMove(m.pos, [2]int{m.pos[0] - 1, m.pos[1]}) {
		if m.nsteps+1 < fewestSteps[m.pos[0]-1][m.pos[1]] {
			nexts = append(nexts, move{pos: [2]int{m.pos[0] - 1, m.pos[1]}, nsteps: m.nsteps + 1})
			fewestSteps[m.pos[0]-1][m.pos[1]] = m.nsteps + 1
		}
	}

	// Check down
	if m.pos[0]+1 < len(height) && canMove(m.pos, [2]int{m.pos[0] + 1, m.pos[1]}) {
		if m.nsteps+1 < fewestSteps[m.pos[0]+1][m.pos[1]] {
			nexts = append(nexts, move{pos: [2]int{m.pos[0] + 1, m.pos[1]}, nsteps: m.nsteps + 1})
			fewestSteps[m.pos[0]+1][m.pos[1]] = m.nsteps + 1
		}
	}

	// Check left
	if m.pos[1]-1 >= 0 && canMove(m.pos, [2]int{m.pos[0], m.pos[1] - 1}) {
		if m.nsteps+1 < fewestSteps[m.pos[0]][m.pos[1]-1] {
			nexts = append(nexts, move{pos: [2]int{m.pos[0], m.pos[1] - 1}, nsteps: m.nsteps + 1})
			fewestSteps[m.pos[0]][m.pos[1]-1] = m.nsteps + 1
		}
	}

	// Check right
	if m.pos[1]+1 < len(height[0]) && canMove(m.pos, [2]int{m.pos[0], m.pos[1] + 1}) {
		if m.nsteps+1 < fewestSteps[m.pos[0]][m.pos[1]+1] {
			nexts = append(nexts, move{pos: [2]int{m.pos[0], m.pos[1] + 1}, nsteps: m.nsteps + 1})
			fewestSteps[m.pos[0]][m.pos[1]+1] = m.nsteps + 1
		}
	}

	return nexts
}

func iniFewestSteps(h, w int, init [2]int) {
	fewestSteps = make([][]int, h)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fewestSteps[i] = append(fewestSteps[i], LIMIT)
		}
	}

	fewestSteps[init[0]][init[1]] = 0
}

func solvePart1(input []string) {
	hm := createHeightMap(input)
	iniFewestSteps(len(hm.height), len(hm.height[0]), hm.start)

	moves := []move{
		{pos: hm.start, nsteps: 0},
	}

	for len(moves) > 0 {
		m := moves[0]
		moves = moves[1:]

		if m.pos[0] == hm.goal[0] && m.pos[1] == hm.goal[1] {
			continue
		}

		nexts := getValidNextMoves(m, hm.height)
		moves = append(moves, nexts...)
	}

	fmt.Println(fewestSteps[hm.goal[0]][hm.goal[1]])

	// for i, row := range fewestSteps {
	// 	for j, n := range row {
	// 		if n < LIMIT {
	// 			fmt.Printf("%c", byte(hm.height[i][j]+'a'))
	// 		} else {
	// 			fmt.Printf(" ")

	// 		}
	// 	}
	// 	fmt.Println()
	// }
}

func main() {
	filename := "input"
	input := getInputFromFile(filename)

	solvePart1(input)
}
