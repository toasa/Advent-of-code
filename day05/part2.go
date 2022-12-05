package main

import (
	"fmt"
	"strconv"
	"strings"
)

func move(src, dst *[]byte, n int) {
	elems := (*src)[len(*src)-n:]
	*src = (*src)[:len(*src)-n]
	*dst = append(*dst, elems...)
}

//         [C] [B] [H]
// [W]     [D] [J] [Q] [B]
// [P] [F] [Z] [F] [B] [L]
// [G] [Z] [N] [P] [J] [S] [V]
// [Z] [C] [H] [Z] [G] [T] [Z]     [C]
// [V] [B] [M] [M] [C] [Q] [C] [G] [H]
// [S] [V] [L] [D] [F] [F] [G] [L] [F]
// [B] [J] [V] [L] [V] [G] [L] [N] [J]
//  1   2   3   4   5   6   7   8   9

func main() {
	stacks := [9][]byte{
		[]byte{'B', 'S', 'V', 'Z', 'G', 'P', 'W'},
		[]byte{'J', 'V', 'B', 'C', 'Z', 'F'},
		[]byte{'V', 'L', 'M', 'H', 'N', 'Z', 'D', 'C'},
		[]byte{'L', 'D', 'M', 'Z', 'P', 'F', 'J', 'B'},
		[]byte{'V', 'F', 'C', 'G', 'J', 'B', 'Q', 'H'},
		[]byte{'G', 'F', 'Q', 'T', 'S', 'L', 'B'},
		[]byte{'L', 'G', 'C', 'Z', 'V'},
		[]byte{'N', 'L', 'G'},
		[]byte{'J', 'F', 'H', 'C'},
	}

	input := `move 5 from 4 to 7
move 8 from 5 to 9
move 6 from 2 to 8
move 7 from 7 to 9
move 1 from 7 to 4
move 2 from 7 to 4
move 9 from 8 to 4
move 16 from 9 to 7
move 1 from 3 to 8
move 15 from 4 to 5
move 3 from 9 to 5
move 2 from 3 to 5
move 1 from 8 to 7
move 3 from 1 to 7
move 5 from 3 to 5
move 13 from 7 to 2
move 5 from 7 to 1
move 7 from 2 to 6
move 2 from 7 to 8
move 3 from 6 to 5
move 2 from 8 to 2
move 2 from 6 to 1
move 11 from 1 to 7
move 2 from 2 to 9
move 8 from 6 to 5
move 2 from 9 to 6
move 3 from 6 to 4
move 1 from 4 to 7
move 22 from 5 to 6
move 13 from 6 to 9
move 5 from 2 to 7
move 6 from 5 to 8
move 13 from 7 to 2
move 2 from 4 to 6
move 5 from 6 to 3
move 2 from 7 to 5
move 3 from 3 to 6
move 2 from 6 to 2
move 8 from 2 to 4
move 2 from 4 to 7
move 2 from 2 to 9
move 5 from 4 to 5
move 2 from 3 to 2
move 1 from 5 to 4
move 6 from 5 to 9
move 1 from 7 to 3
move 1 from 5 to 9
move 5 from 5 to 1
move 1 from 6 to 8
move 1 from 5 to 8
move 4 from 6 to 9
move 8 from 8 to 9
move 1 from 3 to 6
move 4 from 1 to 7
move 3 from 6 to 4
move 7 from 2 to 6
move 27 from 9 to 8
move 3 from 4 to 7
move 6 from 8 to 1
move 1 from 4 to 6
move 1 from 2 to 7
move 7 from 6 to 3
move 1 from 4 to 3
move 4 from 1 to 6
move 1 from 9 to 2
move 1 from 2 to 4
move 1 from 4 to 5
move 3 from 9 to 4
move 5 from 7 to 8
move 2 from 5 to 6
move 4 from 6 to 9
move 10 from 8 to 3
move 2 from 4 to 7
move 3 from 1 to 7
move 2 from 9 to 6
move 6 from 3 to 1
move 7 from 3 to 4
move 2 from 1 to 9
move 4 from 1 to 9
move 1 from 3 to 6
move 1 from 3 to 8
move 2 from 9 to 5
move 2 from 5 to 3
move 3 from 3 to 1
move 1 from 4 to 6
move 5 from 7 to 6
move 2 from 3 to 4
move 2 from 8 to 1
move 9 from 4 to 7
move 4 from 9 to 3
move 2 from 8 to 3
move 1 from 1 to 4
move 1 from 6 to 2
move 1 from 2 to 9
move 6 from 3 to 5
move 2 from 1 to 3
move 1 from 3 to 2
move 1 from 2 to 9
move 8 from 6 to 8
move 2 from 6 to 3
move 1 from 1 to 2
move 7 from 7 to 9
move 13 from 8 to 6
move 1 from 2 to 8
move 6 from 9 to 3
move 1 from 1 to 6
move 2 from 8 to 5
move 5 from 3 to 4
move 2 from 8 to 1
move 8 from 5 to 2
move 4 from 3 to 2
move 5 from 8 to 4
move 2 from 9 to 4
move 4 from 4 to 7
move 10 from 2 to 6
move 1 from 2 to 9
move 24 from 6 to 1
move 17 from 1 to 8
move 1 from 9 to 2
move 2 from 4 to 9
move 10 from 7 to 4
move 1 from 2 to 5
move 5 from 9 to 1
move 1 from 7 to 6
move 12 from 8 to 6
move 1 from 7 to 5
move 2 from 5 to 6
move 16 from 6 to 8
move 12 from 1 to 6
move 2 from 1 to 7
move 9 from 6 to 2
move 2 from 4 to 1
move 1 from 1 to 5
move 7 from 4 to 6
move 13 from 8 to 2
move 5 from 8 to 2
move 2 from 7 to 3
move 2 from 4 to 9
move 1 from 5 to 4
move 3 from 9 to 8
move 2 from 4 to 2
move 2 from 3 to 8
move 1 from 1 to 5
move 1 from 4 to 8
move 6 from 2 to 7
move 1 from 5 to 8
move 1 from 6 to 2
move 7 from 6 to 8
move 1 from 6 to 2
move 24 from 2 to 1
move 10 from 8 to 3
move 4 from 8 to 2
move 4 from 7 to 1
move 5 from 2 to 9
move 1 from 6 to 2
move 10 from 3 to 1
move 2 from 7 to 3
move 2 from 3 to 7
move 2 from 7 to 9
move 35 from 1 to 5
move 28 from 5 to 6
move 2 from 2 to 7
move 19 from 6 to 4
move 3 from 1 to 2
move 3 from 2 to 5
move 23 from 4 to 7
move 2 from 6 to 8
move 4 from 7 to 6
move 3 from 5 to 6
move 13 from 7 to 4
move 2 from 5 to 6
move 2 from 9 to 4
move 5 from 6 to 3
move 6 from 4 to 5
move 1 from 4 to 8
move 4 from 4 to 6
move 5 from 9 to 7
move 2 from 8 to 7
move 5 from 3 to 2
move 4 from 5 to 2
move 5 from 2 to 9
move 4 from 8 to 4
move 1 from 9 to 8
move 2 from 2 to 6
move 4 from 4 to 2
move 3 from 2 to 3
move 3 from 5 to 1
move 2 from 3 to 2
move 3 from 1 to 4
move 1 from 9 to 4
move 5 from 4 to 9
move 2 from 4 to 3
move 5 from 6 to 8
move 1 from 9 to 7
move 2 from 6 to 3
move 1 from 4 to 5
move 1 from 9 to 4
move 6 from 8 to 6
move 2 from 3 to 6
move 2 from 9 to 4
move 2 from 3 to 9
move 1 from 3 to 1
move 17 from 6 to 4
move 1 from 1 to 8
move 1 from 6 to 5
move 1 from 9 to 2
move 11 from 4 to 6
move 9 from 4 to 5
move 7 from 9 to 4
move 2 from 5 to 2
move 1 from 4 to 9
move 5 from 2 to 1
move 1 from 2 to 9
move 4 from 4 to 9
move 4 from 1 to 5
move 1 from 1 to 7
move 1 from 8 to 9
move 8 from 7 to 8
move 4 from 7 to 4
move 9 from 5 to 2
move 2 from 4 to 1
move 11 from 6 to 8
move 2 from 4 to 3
move 2 from 4 to 8
move 1 from 1 to 4
move 3 from 2 to 8
move 1 from 1 to 3
move 3 from 3 to 9
move 8 from 9 to 6
move 1 from 4 to 8
move 2 from 9 to 3
move 5 from 6 to 9
move 7 from 5 to 6
move 2 from 3 to 4
move 5 from 7 to 9
move 2 from 4 to 5
move 2 from 2 to 3
move 10 from 9 to 5
move 2 from 6 to 3
move 6 from 2 to 7
move 10 from 5 to 3
move 6 from 7 to 1
move 2 from 1 to 7
move 4 from 3 to 9
move 3 from 8 to 2
move 2 from 7 to 5
move 19 from 8 to 7
move 4 from 5 to 9
move 4 from 9 to 8
move 1 from 2 to 5
move 3 from 6 to 8
move 1 from 5 to 9
move 5 from 9 to 7
move 6 from 3 to 8
move 1 from 3 to 8
move 2 from 3 to 2
move 23 from 7 to 6
move 10 from 8 to 4
move 4 from 4 to 9
move 4 from 2 to 6
move 1 from 3 to 8
move 4 from 8 to 4
move 31 from 6 to 4
move 9 from 4 to 5
move 8 from 5 to 3
move 1 from 6 to 7
move 2 from 5 to 7
move 4 from 9 to 2
move 21 from 4 to 8
move 4 from 2 to 9
move 3 from 3 to 9
move 2 from 7 to 9
move 11 from 4 to 9
move 1 from 8 to 5
move 1 from 5 to 9
move 9 from 9 to 3
move 3 from 1 to 5
move 2 from 5 to 8
move 11 from 3 to 6
move 4 from 6 to 3
move 2 from 8 to 3
move 10 from 9 to 6
move 22 from 8 to 9
move 1 from 1 to 8
move 4 from 6 to 3
move 2 from 7 to 6
move 3 from 8 to 3
move 14 from 3 to 2
move 1 from 3 to 4
move 1 from 2 to 4
move 2 from 9 to 1
move 1 from 5 to 7
move 1 from 3 to 2
move 14 from 6 to 5
move 13 from 5 to 2
move 1 from 5 to 6
move 1 from 7 to 9
move 8 from 9 to 4
move 2 from 6 to 7
move 23 from 2 to 4
move 2 from 1 to 4
move 2 from 2 to 5
move 1 from 5 to 1
move 1 from 7 to 2
move 1 from 5 to 9
move 16 from 9 to 5
move 1 from 2 to 4
move 13 from 5 to 3
move 1 from 1 to 4
move 1 from 7 to 1
move 1 from 5 to 3
move 2 from 5 to 7
move 2 from 7 to 1
move 9 from 3 to 2
move 2 from 1 to 7
move 1 from 1 to 9
move 19 from 4 to 2
move 1 from 9 to 7
move 1 from 7 to 8
move 23 from 2 to 8
move 2 from 7 to 2
move 12 from 4 to 5
move 12 from 5 to 1
move 5 from 2 to 9
move 2 from 2 to 7
move 5 from 8 to 1
move 3 from 9 to 4
move 1 from 2 to 8
move 1 from 2 to 4
move 4 from 8 to 1
move 2 from 3 to 1
move 2 from 7 to 5
move 1 from 4 to 9
move 8 from 4 to 7
move 13 from 8 to 6
move 1 from 3 to 1
move 13 from 6 to 7
move 13 from 7 to 6
move 7 from 1 to 4
move 5 from 7 to 3
move 3 from 4 to 3
move 13 from 6 to 1
move 3 from 8 to 6
move 8 from 3 to 8
move 12 from 1 to 8
move 1 from 3 to 5
move 6 from 1 to 7
move 3 from 6 to 8
move 1 from 3 to 8
move 1 from 9 to 2
move 3 from 5 to 6
move 1 from 7 to 3
move 8 from 7 to 1
move 2 from 6 to 2
move 3 from 4 to 3
move 2 from 9 to 2
move 6 from 8 to 9
move 5 from 2 to 5
move 2 from 3 to 4
move 5 from 5 to 4
move 1 from 3 to 9
move 8 from 4 to 5
move 1 from 6 to 8
move 2 from 1 to 4
move 1 from 1 to 4
move 3 from 1 to 5
move 3 from 1 to 6
move 7 from 1 to 9
move 2 from 6 to 9
move 1 from 3 to 5
move 17 from 8 to 7
move 17 from 7 to 6
move 5 from 5 to 2
move 5 from 2 to 1
move 13 from 6 to 2
move 1 from 1 to 4
move 5 from 5 to 1
move 1 from 1 to 5
move 10 from 9 to 1
move 13 from 1 to 8
move 13 from 8 to 4
move 5 from 6 to 7
move 8 from 1 to 7
move 1 from 1 to 3
move 12 from 2 to 6
move 1 from 3 to 8
move 6 from 6 to 2
move 2 from 5 to 1
move 5 from 2 to 5
move 2 from 5 to 9
move 12 from 4 to 2
move 1 from 6 to 2
move 15 from 2 to 1
move 1 from 8 to 6
move 2 from 7 to 3
move 2 from 4 to 2
move 1 from 2 to 9
move 1 from 2 to 6
move 7 from 7 to 3
move 1 from 4 to 1
move 17 from 1 to 2
move 3 from 6 to 4
move 1 from 3 to 8
move 3 from 9 to 6
move 4 from 6 to 3
move 13 from 2 to 9
move 3 from 2 to 8
move 2 from 5 to 1
move 6 from 8 to 2
move 1 from 6 to 2
move 3 from 2 to 7
move 3 from 1 to 6
move 2 from 9 to 8
move 6 from 9 to 8
move 8 from 9 to 3
move 7 from 7 to 4
move 20 from 3 to 7
move 4 from 6 to 8
move 1 from 8 to 6
move 2 from 6 to 4
move 3 from 2 to 1
move 2 from 9 to 6
move 9 from 8 to 6
move 3 from 1 to 9
move 9 from 4 to 8
move 1 from 5 to 6
move 3 from 4 to 2
move 1 from 5 to 3
move 8 from 6 to 4
move 4 from 9 to 3
move 10 from 8 to 6
move 5 from 2 to 3
move 3 from 6 to 4
move 10 from 3 to 1
move 11 from 4 to 1
move 1 from 8 to 2
move 2 from 4 to 2
move 1 from 4 to 9
move 10 from 6 to 3
move 21 from 1 to 5
move 2 from 2 to 7
move 1 from 9 to 6
move 1 from 6 to 3
move 1 from 6 to 7
move 11 from 5 to 6
move 1 from 2 to 8
move 1 from 5 to 9
move 11 from 6 to 3
move 1 from 8 to 4
move 1 from 4 to 1
move 3 from 5 to 7
move 1 from 1 to 5
move 5 from 5 to 8
move 23 from 7 to 9
move 5 from 8 to 4
move 1 from 5 to 2
move 12 from 3 to 4
move 6 from 3 to 6
move 1 from 5 to 2
move 8 from 9 to 2
move 1 from 7 to 8
move 2 from 7 to 9
move 4 from 3 to 5
move 1 from 5 to 9
move 1 from 6 to 5
move 4 from 6 to 5
move 3 from 2 to 1
move 3 from 1 to 3
move 8 from 9 to 1
move 4 from 2 to 9
move 1 from 9 to 7
move 14 from 4 to 8
move 3 from 3 to 4
move 1 from 5 to 8
move 2 from 8 to 6
move 2 from 6 to 7
move 4 from 4 to 3
move 12 from 9 to 1
move 1 from 3 to 2
move 6 from 8 to 2
move 1 from 7 to 1
move 5 from 2 to 3
move 21 from 1 to 3
move 5 from 5 to 4
move 1 from 8 to 5
move 2 from 2 to 7
move 1 from 6 to 1
move 2 from 9 to 2
move 1 from 2 to 9
move 1 from 1 to 5
move 4 from 3 to 5
move 7 from 8 to 1
move 6 from 1 to 9
move 1 from 2 to 5
move 6 from 9 to 7
move 8 from 3 to 4
move 2 from 4 to 8
move 1 from 1 to 6
move 10 from 3 to 9
move 12 from 4 to 2
move 1 from 8 to 1
`

	moves := []string{}
	for {
		i := strings.Index(input, "\n")
		if i < 0 {
			break
		}
		moves = append(moves, input[:i])

		// Skip '\n'
		input = input[i+1:]
	}

	for _, m := range moves {
		m = m[len("move "):]
		i := strings.Index(m, " ")
		n, _ := strconv.Atoi(m[:i])

		m = m[i+1+len("from "):]
		i = strings.Index(m, " ")
		src, _ := strconv.Atoi(m[:i])

		m = m[i+1+len("to "):]
		dst, _ := strconv.Atoi(m[:i])

		// fmt.Println("src:", src, "dst:", dst, "n:", n)

		move(&stacks[src-1], &stacks[dst-1], n)
	}

	for _, stack := range stacks {
		fmt.Printf("%c", stack[len(stack)-1])
	}
	fmt.Println("")
}
