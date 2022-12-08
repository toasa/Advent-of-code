package main

import (
	"bufio"
	"fmt"
	"os"
)

func splitIntoLines(filename string) []string {
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

const LEN = 99

func getTrees(lines []string) [LEN][LEN]int {
	trees := [LEN][LEN]int{}
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			c := line[j]
			trees[i][j] = int(c - '0')
		}
	}
	return trees
}

func calcScenicScore(trees [LEN][LEN]int, i, j int) int {
	h := trees[i][j]

	visibleTreesUp := 0
	visibleTreesDown := 0
	visibleTreesLeft := 0
	visibleTreesRight := 0

	doneUp := false
	doneDown := false
	doneLeft := false
	doneRight := false

	for n := 1; n < LEN; n++ {
		if i+n < LEN && !doneUp {
			visibleTreesUp++
			if trees[i+n][j] >= h {
				doneUp = true
			}
		}

		if i-n >= 0 && !doneDown {
			visibleTreesDown++
			if trees[i-n][j] >= h {
				doneDown = true
			}
		}

		if j+n < LEN && !doneRight {
			visibleTreesRight++
			if trees[i][j+n] >= h {
				doneRight = true
			}
		}

		if j-n >= 0 && !doneLeft {
			visibleTreesLeft++
			if trees[i][j-n] >= h {
				doneLeft = true
			}
		}
	}

	return visibleTreesDown * visibleTreesLeft * visibleTreesRight * visibleTreesUp
}

func solvePart2(lines []string) {
	trees := getTrees(lines)

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	score := 0
	for i := 1; i < LEN-1; i++ {
		for j := 1; j < LEN-1; j++ {
			score = max(score, calcScenicScore(trees, i, j))
		}
	}

	fmt.Println(score)
}

func main() {
	filename := "input"
	lines := splitIntoLines(filename)

	solvePart2(lines)
}
