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

func solvePart1(lines []string) {
	trees := getTrees(lines)

	visibleFromRight := [LEN][LEN]bool{}
	visibleFromLeft := [LEN][LEN]bool{}
	visibleFromTop := [LEN][LEN]bool{}
	visibleFromBottom := [LEN][LEN]bool{}

	// Every tree of edge are visible
	for i := 0; i < LEN; i++ {
		visibleFromTop[0][i] = true
		visibleFromBottom[LEN-1][i] = true
		visibleFromLeft[i][0] = true
		visibleFromRight[i][LEN-1] = true
	}

	// Check view from the left
	for i := 0; i < LEN; i++ {
		h := trees[i][0]
		for j := 1; j < LEN; j++ {
			if trees[i][j] > h {
				h = trees[i][j]
				visibleFromLeft[i][j] = true
			}
		}
	}

	// Check view from the right
	for i := 0; i < LEN; i++ {
		h := trees[i][LEN-1]
		for j := LEN - 2; j >= 0; j-- {
			if trees[i][j] > h {
				h = trees[i][j]
				visibleFromRight[i][j] = true
			}
		}
	}

	// Check view from the top
	for j := 0; j < LEN; j++ {
		h := trees[0][j]
		for i := 1; i < LEN; i++ {
			if trees[i][j] > h {
				h = trees[i][j]
				visibleFromTop[i][j] = true
			}
		}
	}

	// Check view from the bottom
	for j := 0; j < LEN; j++ {
		h := trees[LEN-1][j]
		for i := LEN - 2; i >= 0; i-- {
			if trees[i][j] > h {
				h = trees[i][j]
				visibleFromBottom[i][j] = true
			}
		}
	}

	count := LEN*2 + (LEN-2)*2
	for i := 1; i < LEN-1; i++ {
		for j := 1; j < LEN-1; j++ {
			v := visibleFromBottom[i][j] || visibleFromTop[i][j] ||
				visibleFromLeft[i][j] || visibleFromRight[i][j]

			if v {
				count++
			}
		}
	}

	fmt.Println(count)
}

func main() {
	filename := "input"
	lines := splitIntoLines(filename)

	solvePart1(lines)
}
