package day10

import (
	"fmt"
	"log"

	"theelements.org/advent-of-code/common"
)

type loc struct {
	height int
	peaks  map[string]any
	count  int
}

func readBoard(input []string) ([][]*loc, []*loc) {
	board := make([][]*loc, len(input))
	trailheads := make([]*loc, 0)
	for i, line := range input {
		board[i] = make([]*loc, len(line))
		for j, r := range line {
			height := common.Atoi(string(r))
			var peaks map[string]any
			if height == 0 {
				peaks = make(map[string]any)
			}
			board[i][j] = &loc{height: height, peaks: peaks, count: 0}
			if board[i][j].height == 0 {
				trailheads = append(trailheads, board[i][j])
			}
		}
	}
	return board, trailheads
}

func dfs(board [][]*loc, row, col, startRow, startCol int) {
	if board[row][col].height == 0 {
		k := fmt.Sprintf("(%d, %d)", startRow, startCol)
		board[row][col].peaks[k] = true
		board[row][col].count++
		return
	}

	// maxRow and maxCol, for bounds checking
	mr := len(board)
	mc := len(board[0])

	// The height to look for
	h := board[row][col].height - 1

	var dirs = []struct{ r, c int }{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}
	for _, d := range dirs {
		// newRow, newCol
		nr := row + d.r
		nc := col + d.c
		if nr < 0 || nr >= mr || nc < 0 || nc >= mc {
			continue
		}
		if board[nr][nc].height == h {
			dfs(board, nr, nc, startRow, startCol)
		}
	}
}

func P1() {
	input := common.ReadFile("./2024/day10/input.txt")
	board, trailheads := readBoard(input)
	score := P1Solution(board, trailheads)
	log.Printf("The trailhead score is: %d", score)
}

// w and h are the max width and heights so we can determine
// if an antinode is outside of the map
func P1Solution(board [][]*loc, trailheads []*loc) int {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col].height == 9 {
				dfs(board, row, col, row, col)
			}
		}
	}

	sum := 0
	for _, th := range trailheads {
		sum += len(th.peaks)
	}
	return sum
}

func P2() {
	input := common.ReadFile("./2024/day10/input.txt")
	board, trailheads := readBoard(input)
	score := P2Solution(board, trailheads)
	log.Printf("The trailhead score is: %d", score)
}

func P2Solution(board [][]*loc, trailheads []*loc) int {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col].height == 9 {
				dfs(board, row, col, row, col)
			}
		}
	}

	sum := 0
	for _, th := range trailheads {
		sum += th.count
	}
	return sum
}
