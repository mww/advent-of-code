package day04

import (
	"log"

	"theelements.org/advent-of-code/common"
)

func P1() {
	input := common.ReadFile("./2024/day04/input.txt")
	board := createBoard(input)
	count := P1Solution(board)
	log.Printf("XMAS count is: %d", count)
}

func P1Solution(board [][]rune) int64 {
	height := len(board)
	width := len(board[0])
	count := int64(0)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if board[x][y] == 'X' {
				if search(x, y, 1, 0, board) {
					count++
				}
				if search(x, y, -1, 0, board) {
					count++
				}
				if search(x, y, 0, 1, board) {
					count++
				}
				if search(x, y, 0, -1, board) {
					count++
				}
				if search(x, y, 1, 1, board) {
					count++
				}
				if search(x, y, 1, -1, board) {
					count++
				}
				if search(x, y, -1, 1, board) {
					count++
				}
				if search(x, y, -1, -1, board) {
					count++
				}
			}
		}
	}
	return count
}

func P2() {
	input := common.ReadFile("./2024/day04/input.txt")
	board := createBoard(input)
	count := P2Solution(board)
	log.Printf("X-MAS count is: %d", count)
}

func P2Solution(board [][]rune) int64 {
	height := len(board)
	width := len(board[0])
	count := int64(0)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if board[x][y] == 'A' {
				if searchPart2(x, y, board) {
					count++
				}
			}
		}
	}
	return count
}

func createBoard(input []string) [][]rune {
	height := len(input)
	width := len(input[0])

	results := make([][]rune, 0, height)
	for _, l := range input {
		line := make([]rune, width)
		for i, r := range l {
			line[i] = r
		}
		results = append(results, line)
	}
	return results
}

func search(x, y, stepX, stepY int, board [][]rune) bool {
	h := len(board)
	w := len(board[0])

	for _, r := range "XMAS" {
		if x < 0 || x >= w {
			return false
		}
		if y < 0 || y >= h {
			return false
		}
		if board[x][y] != r {
			return false
		}
		x += stepX
		y += stepY
	}
	return true
}

// We will get here starting with an "A" since it is always the middle of all
// of the X-MAS's. Then look for the possible combinations that make up the
// pattern we are looking for.
func searchPart2(x, y int, board [][]rune) bool {
	h := len(board)
	w := len(board[0])

	// Make sure we aren't on an edge
	if x < 1 || x > w-2 {
		return false
	}
	if y < 1 || y > h-2 {
		return false
	}

	if board[x][y] != 'A' {
		return false
	}
	left := (board[x-1][y-1] == 'M' && board[x+1][y+1] == 'S') || (board[x-1][y-1] == 'S' && board[x+1][y+1] == 'M')
	right := (board[x-1][y+1] == 'M' && board[x+1][y-1] == 'S') || (board[x-1][y+1] == 'S' && board[x+1][y-1] == 'M')

	return left && right
}
