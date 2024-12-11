package day08

import (
	"reflect"
	"testing"
)

var input = []string{
	"............",
	"........0...",
	".....0......",
	".......0....",
	"....0.......",
	"......A.....",
	"............",
	"............",
	"........A...",
	".........A..",
	"............",
	"............",
}

func TestP1Solution(t *testing.T) {
	data := readData(input)
	n := P1Solution(12, 12, data)
	if n != 14 {
		t.Errorf("number of antinodes should be 14, but was: %d", n)
	}
}

func TestP2Solution(t *testing.T) {
	data := readData(input)
	n := P2Solution(12, 12, data)
	if n != 34 {
		t.Errorf("number of antinodes should be 34, but was: %d", n)
	}
}

func TestReadData(t *testing.T) {
	expected := map[rune][]loc{
		'A': {{5, 6}, {8, 8}, {9, 9}},
		'0': {{1, 8}, {2, 5}, {3, 7}, {4, 4}},
	}
	data := readData(input)

	if !reflect.DeepEqual(expected, data) {
		t.Errorf("input data not as expected: %v", data)
	}
}
