package day10

import (
	"testing"
)

var input = []string{
	"89010123",
	"78121874",
	"87430965",
	"96549874",
	"45678903",
	"32019012",
	"01329801",
	"10456732",
}

func TestP1Solution(t *testing.T) {
	board, th := readBoard(input)
	s := P1Solution(board, th)
	if s != 36 {
		t.Errorf("expected score of threadheads to be 36, but was %d", s)
	}
}

func TestP2Solution(t *testing.T) {
	board, th := readBoard(input)
	s := P2Solution(board, th)
	if s != 81 {
		t.Errorf("expected score of threadheads to be 81, but was %d", s)
	}
}
