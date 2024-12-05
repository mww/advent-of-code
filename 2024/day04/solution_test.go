package day04

import (
	"testing"
)

var input = []string{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func TestP1Solution(t *testing.T) {
	board := createBoard(input)
	count := P1Solution(board)
	if count != 18 {
		t.Errorf("expected to find XMAS 18 times, but found it: %d", count)
	}
}

func TestP2Solution(t *testing.T) {
	board := createBoard(input)
	count := P2Solution(board)
	if count != 9 {
		t.Errorf("expected to find X-MAS 9 times, but found it: %d", count)
	}
}
