package day06

import (
	"fmt"
	"log"

	"theelements.org/advent-of-code/common"
)

func P1() {
	input := common.ReadFile("./2024/day06/input.txt")
	board := buildBoard(input)

	visited := P1Solution(board)
	log.Printf("Unique squares the guard visited: %d", visited)
}

// Return -1 if we think the guard is in an infinate loop
func P1Solution(board [][]rune) int {
	// Make a copy of the board so we don't modify the original
	b := copyBoard(board)

	deltas := map[rune]struct {
		dx, dy int
	}{
		'N': {dx: 0, dy: -1},
		'S': {dx: 0, dy: 1},
		'E': {dx: 1, dy: 0},
		'W': {dx: -1, dy: 0},
	}

	boardHeight := len(board)
	boardWidth := len(board[0])

	visitedMap := make(map[string]bool)

	guardX, guardY, guardDir := findGuard(b)
	for {
		if guardX < 0 || guardY < 0 || guardX >= boardWidth || guardY >= boardHeight {
			// Guard has walked off the board
			break
		}
		// Mark the guard's current position so we can come back and could it later
		b[guardY][guardX] = 'X'

		key := fmt.Sprintf("%d-%d-%c", guardY, guardX, guardDir)
		if _, ok := visitedMap[key]; ok {
			return -1
		} else {
			visitedMap[key] = true
		}

		delta, ok := deltas[guardDir]
		if !ok {
			log.Fatalf("guard not facing a normal direction: '%v'", guardDir)
		}

		if guardX+delta.dx < 0 || guardY+delta.dy < 0 || guardX+delta.dx >= boardWidth || guardY+delta.dy >= boardHeight {
			// Check if the next move will take us off the board
			break
		}

		// Account for double turn
		for {
			// There is an obstacle, have the guard turn right
			if b[guardY+delta.dy][guardX+delta.dx] == '#' {
				switch guardDir {
				case 'N':
					guardDir = 'E'
				case 'E':
					guardDir = 'S'
				case 'S':
					guardDir = 'W'
				case 'W':
					guardDir = 'N'
				}

				delta = deltas[guardDir]
			} else {
				break
			}
		}

		guardX = guardX + delta.dx
		guardY = guardY + delta.dy
	}

	// Count the visited locations
	visited := 0
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == 'X' {
				visited++
			}
		}
	}
	return visited
}

func P2() {
	input := common.ReadFile("./2024/day06/input.txt")
	board := buildBoard(input)

	count := P2Solution(board)
	log.Printf("Number of obstacle placement options: %d", count)
}

func P2Solution(board [][]rune) int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '^' {
				//log.Printf("adding obstacle at (%d,%d)", j, i)
				b := copyBoard(board)
				b[i][j] = '#' // add a new obstacle

				if P1Solution(b) == -1 {
					//log.Printf("found a loop at (%d,%d)", j, i)
					// for _, l := range b {
					// 	log.Printf("%v", l)
					// }
					count++
				}
			} else {
				//.Printf("skipping obstacle at (%d,%d)", j, i)
			}
		}
	}
	return count
}

func copyBoard(board [][]rune) [][]rune {
	b := make([][]rune, len(board))
	for i, row := range board {
		b[i] = make([]rune, len(row))
		copy(b[i], row)
	}
	return b
}

func buildBoard(input []string) [][]rune {
	board := make([][]rune, len(input))
	for i, line := range input {
		row := make([]rune, len(line))
		for j, c := range line {
			row[j] = c
		}
		board[i] = row
	}
	return board
}

func findGuard(board [][]rune) (int, int, rune) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '^' {
				return j, i, 'N'
			} else if board[i][j] == '>' {
				return j, i, 'E'
			} else if board[i][j] == 'V' {
				return j, i, 'S'
			} else if board[i][j] == '<' {
				return j, i, 'W'
			}
		}
	}

	log.Fatal("did not initialize the guard")
	return -1, -1, 'X'
}
