package day06

import (
	"testing"
)

var input []string = []string{
	"....#.....",
	".........#",
	"..........",
	"..#.......",
	".......#..",
	"..........",
	".#..^.....",
	"........#.",
	"#.........",
	"......#...",
}

func TestP1Solution(t *testing.T) {
	board := buildBoard(input)
	visited := P1Solution(board)
	if visited != 41 {
		t.Errorf("expected to visit 41 locations, but was: %d", visited)
	}
}

func TestP2Solution(t *testing.T) {
	board := buildBoard(input)
	count := P2Solution(board)
	if count != 6 {
		t.Errorf("expected 6 options, but was: %d", count)
	}
}

func TestFindGuard(t *testing.T) {
	board := buildBoard(input)
	guardX, guardY, guardDir := findGuard(board)
	if guardX != 4 {
		t.Errorf("guardX is wrong: %d", guardX)
	}
	if guardY != 6 {
		t.Errorf("guardY is wrong: %d", guardY)
	}
	if guardDir != 'N' {
		t.Errorf("guardDir is wrong: %v", guardDir)
	}
}
