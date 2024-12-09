package day07

import (
	"reflect"
	"slices"
	"testing"
)

var input = []string{
	"190: 10 19",
	"3267: 81 40 27",
	"83: 17 5",
	"156: 15 6",
	"7290: 6 8 6 15",
	"161011: 16 10 13",
	"192: 17 8 14",
	"21037: 9 7 18 13",
	"292: 11 6 16 20",
}

func TestP1Solution(t *testing.T) {
	s := P1Solution(input)
	if s != 3749 {
		t.Errorf("expected 3749 but got: %d", s)
	}
}

func TestP2Solution(t *testing.T) {
	s := P2Solution(input)
	if s != 11387 {
		t.Errorf("expected 11387 but got: %d", s)
	}
}

func TestGenerateOpsOptions(t *testing.T) {
	possOps := []rune{'*', '+'}

	var testCases = []struct {
		num int
		e   [][]rune
	}{
		{num: 1, e: [][]rune{{'*'}, {'+'}}},
		{num: 2, e: [][]rune{{'*', '*'}, {'*', '+'}, {'+', '*'}, {'+', '+'}}},
		{num: 3, e: [][]rune{{'*', '*', '*'}, {'*', '*', '+'}, {'*', '+', '*'},
			{'*', '+', '+'}, {'+', '*', '*'}, {'+', '*', '+'}, {'+', '+', '*'}, {'+', '+', '+'}}},
	}
	for _, tc := range testCases {
		acc := make([][]rune, 0)
		cur := make([]rune, 0)
		acc = generateOpsOptions(acc, cur, tc.num, possOps)
		if !reflect.DeepEqual(acc, tc.e) {
			t.Errorf("expected '%v' but got '%v'", tc.e, acc)
		}
	}
}

func TestGenerateOpsOptionsPart2(t *testing.T) {
	possOps := []rune{'*', '+', '|'}

	var testCases = []struct {
		num int
		e   [][]rune
	}{
		{num: 1, e: [][]rune{{'*'}, {'+'}, {'|'}}},
		{num: 2, e: [][]rune{
			{'*', '*'}, {'*', '+'}, {'*', '|'},
			{'+', '*'}, {'+', '+'}, {'+', '|'},
			{'|', '*'}, {'|', '+'}, {'|', '|'}}},
	}
	for _, tc := range testCases {
		acc := make([][]rune, 0)
		cur := make([]rune, 0)
		acc = generateOpsOptions(acc, cur, tc.num, possOps)
		if !reflect.DeepEqual(acc, tc.e) {
			t.Errorf("expected '%v' but got '%v'", tc.e, acc)
		}
	}
}

func TestCouldBeTrue(t *testing.T) {
	var testCases = []struct {
		sum    int64
		values []int64
		poss   bool
		exOps  []rune
	}{
		{sum: 190, values: []int64{10, 19}, poss: true, exOps: []rune{'*'}},
		{sum: 3267, values: []int64{81, 40, 27}, poss: true, exOps: []rune{'*', '+'}},
		{sum: 83, values: []int64{17, 5}, poss: false, exOps: nil},
		{sum: 156, values: []int64{15, 6}, poss: false, exOps: nil},
		{sum: 7290, values: []int64{6, 8, 6, 15}, poss: false, exOps: nil},
		{sum: 161011, values: []int64{16, 10, 13}, poss: false, exOps: nil},
		{sum: 192, values: []int64{17, 8, 14}, poss: false, exOps: nil},
		{sum: 21037, values: []int64{9, 7, 18, 13}, poss: false, exOps: nil},
		{sum: 292, values: []int64{11, 6, 16, 20}, poss: true, exOps: []rune{'+', '*', '+'}},
	}

	possOps := []rune{'*', '+'}

	for _, tc := range testCases {
		res, ops := couldBeTrue(tc.sum, tc.values, possOps)
		if res != tc.poss {
			t.Errorf("expected case to be '%v', but was '%v'", tc.poss, res)
		}

		var eq = func(r1, r2 rune) bool {
			return r1 == r2
		}

		if !slices.EqualFunc(ops, tc.exOps, eq) {
			t.Errorf("expected ops to be '%v', but were '%v'", tc.exOps, ops)
		}
	}
}

func TestCouldBeTrueP2(t *testing.T) {
	var testCases = []struct {
		sum    int64
		values []int64
		poss   bool
		exOps  []rune
	}{
		{sum: 190, values: []int64{10, 19}, poss: true, exOps: []rune{'*'}},
		{sum: 3267, values: []int64{81, 40, 27}, poss: true, exOps: []rune{'*', '+'}},
		{sum: 83, values: []int64{17, 5}, poss: false, exOps: nil},
		{sum: 156, values: []int64{15, 6}, poss: true, exOps: []rune{'|'}},
		{sum: 7290, values: []int64{6, 8, 6, 15}, poss: true, exOps: []rune{'*', '|', '*'}},
		{sum: 161011, values: []int64{16, 10, 13}, poss: false, exOps: nil},
		{sum: 192, values: []int64{17, 8, 14}, poss: true, exOps: []rune{'|', '+'}},
		{sum: 21037, values: []int64{9, 7, 18, 13}, poss: false, exOps: nil},
		{sum: 292, values: []int64{11, 6, 16, 20}, poss: true, exOps: []rune{'+', '*', '+'}},
		{sum: 101112, values: []int64{10, 11, 12}, poss: true, exOps: []rune{'|', '|'}},
	}

	possOps := []rune{'*', '+', '|'}

	for _, tc := range testCases {
		res, ops := couldBeTrue(tc.sum, tc.values, possOps)
		if res != tc.poss {
			t.Errorf("expected case to be '%v', but was '%v'", tc.poss, res)
		}

		var eq = func(r1, r2 rune) bool {
			return r1 == r2
		}

		if !slices.EqualFunc(ops, tc.exOps, eq) {
			t.Errorf("expected ops to be '%v', but were '%v'", tc.exOps, ops)
		}
	}
}
