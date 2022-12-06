package day04

import (
	"fmt"
	"log"
	"strings"

	"theelements.org/advent-of-code/common"
)

func P1() {
	input, boards := readInput("./2021/day04/input.txt")

	for _, i := range input {
		for _, b := range boards {
			s := b.mark(i)
			if s != -1 {
				fmt.Printf("board:\n%v\nscore: %d\n", b, s)
				return
			}
		}
	}
}

func P2() {
	input, boards := readInput("./2021/day04/input.txt")

	for _, i := range input {
		next := make([]*board, 0)
		// fmt.Printf("for iteration %d, there are %d boards\n", x, len(boards))
		for _, b := range boards {
			s := b.mark(i)

			if len(boards) == 1 {
				// We are just waiting for the last board to be a winner
				if s != -1 {
					fmt.Printf("last board:\n%v\nscore: %d\n", b, s)
					return
				}
				next = append(next, b)
			} else {
				if s == -1 {
					next = append(next, b)
				} else {
					// fmt.Printf("removing winning board:\n%v\n", b)
				}
			}
		}
		boards = next
	}
}

func readInput(f string) ([]int, []*board) {
	input := make([]int, 0)
	boards := make([]*board, 0)

	lines := common.ReadFile(f)

	parts := strings.Split(lines[0], ",")
	for _, p := range parts {
		i := common.Atoi(p)
		input = append(input, i)
	}

	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		var b board
		row := 0
		for {
			line = lines[i]
			i++

			parts := strings.Fields(line)
			if len(parts) != 5 {
				log.Fatalf("wrong number of parts after split: '%s', %d", line, len(parts))
			}
			for col, p := range parts {
				x := common.Atoi(p)
				b.b[row][col] = square{num: x, marked: false}
			}

			row++
			if row == 5 {
				break
			}
		}

		//log.Printf("board:\n%v", &b)
		boards = append(boards, &b)
	}

	return input, boards
}

type square struct {
	num    int
	marked bool
}

type board struct {
	b [5][5]square
}

func (b *board) mark(x int) int {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.b[i][j].num == x {
				b.b[i][j].marked = true
				break
			}
		}
	}

	// Now check if the board is a winner
	// rows first
	for i := 0; i < 5; i++ {
		inARow := 0
		for j := 0; j < 5; j++ {
			if b.b[i][j].marked {
				inARow++
			} else {
				break
			}
		}
		if inARow == 5 {
			return b.score(x)
		}
	}

	// now columns
	for j := 0; j < 5; j++ {
		inACol := 0
		for i := 0; i < 5; i++ {
			if b.b[i][j].marked {
				inACol++
			} else {
				break
			}
		}

		if inACol == 5 {
			return b.score(x)
		}
	}

	return -1
}

func (b *board) score(x int) int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.b[i][j].marked {
				score += b.b[i][j].num
			}
		}
	}
	return score * x
}

func (b *board) String() string {
	var x strings.Builder
	for _, r := range b.b {
		x.WriteString(fmt.Sprintf("%2d %2d %2d %2d %2d\n", r[0].num, r[1].num, r[2].num, r[3].num, r[4].num))
	}
	return x.String()
}
