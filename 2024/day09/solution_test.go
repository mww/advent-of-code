package day09

import (
	"reflect"
	"testing"
)

var input = "2333133121414131402"

func TestMakeDataSlice(t *testing.T) {
	data := makeDataSlice(input)
	expected := []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9}

	if !reflect.DeepEqual(expected, data) {
		t.Errorf("data slice not expected, got: %v", expected)
	}
}

func TestP1Solution(t *testing.T) {
	data := makeDataSlice(input)
	cs := P1Solution(data)
	if cs != 1928 {
		t.Errorf("the checksum should be 1928, but was: %d", cs)
	}
}

func TestP2Solution(t *testing.T) {
	data := makeDataSlice(input)
	cs := P2Solution(data)
	if cs != 2858 {
		t.Errorf("the checksum should be 2858, but was: %d", cs)
	}
}
