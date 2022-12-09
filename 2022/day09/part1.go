package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDist(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func getNextTailMove(hX, hY, tX, tY int) (int, int) {
	// Overlapping case
	if hX == tX && hY == tY {
		return tX, tY
	}

	md := manhattanDist(hX, hY, tX, tY)
	if md <= 2 {
		// Adjacent case
		if md == 1 {
			return tX, tY
		}

		// Diagonally adjacent case
		if hX != tX && hY != tY {
			return tX, tY
		}

		if hX == tX {
			return tX, (hY + tY) / 2
		}

		return (tX + hX) / 2, tY
	}

	dX, dY := abs(hX-tX), abs(hY-tY)
	if dX > dY {
		return (hX + tX) / 2, hY
	} else if dX < dY {
		return hX, (hY + tY) / 2
	}

	panic("unreachable")
}

const LEN = 1000

func solvePart1(moves []byte, dists []int) {
	visited := [LEN][LEN]bool{}

	hX, hY := LEN/2, LEN/2
	tX, tY := hX, hY
	visited[tX][tY] = true

	for i, move := range moves {
		dist := dists[i]

		// fmt.Printf("=== move %c, %d times ===\n", move, dist)

		for j := 0; j < dist; j++ {
			switch move {
			case 'U':
				hY += 1
			case 'D':
				hY -= 1
			case 'L':
				hX -= 1
			case 'R':
				hX += 1
			default:
				panic("Invalid move")
			}

			tX, tY = getNextTailMove(hX, hY, tX, tY)

			// fmt.Printf("Head: (%d, %d), Tail: (%d, %d)\n", hX, hY, tX, tY)

			visited[tX][tY] = true
		}
	}

	count := 0
	for i := 0; i < LEN; i++ {
		for j := 0; j < LEN; j++ {
			if visited[i][j] {
				count += 1
			}
		}
	}
	fmt.Println(count)
}

func main() {
	filename := "input"
	input := getInputFromFile(filename)

	moves := make([]byte, len(input))
	dist := make([]int, len(input))

	for i, line := range input {
		moves[i] = line[0]
		n, _ := strconv.Atoi(line[2:])
		dist[i] = n
	}

	solvePart1(moves, dist)
}
