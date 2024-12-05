package day05

import (
	"reflect"
	"testing"

	"theelements.org/advent-of-code/common"
)

var rules = map[int]map[int]any{
	29: {13: true},
	47: {13: true, 29: true, 53: true, 61: true},
	53: {13: true, 29: true},
	61: {13: true, 29: true, 53: true},
	75: {13: true, 29: true, 47: true, 53: true, 61: true},
	97: {13: true, 29: true, 47: true, 53: true, 61: true, 75: true},
}

var expectedValid = [][]int{
	{75, 47, 61, 53, 29},
	{97, 61, 53, 29, 13},
	{75, 29, 13},
}

var expectedInvalid = [][]int{
	{75, 97, 47, 61, 53},
	{61, 13, 29},
	{97, 13, 75, 29, 47},
}

func TestP1Solution(t *testing.T) {
	pageLists := make([][]int, 0)
	pageLists = append(pageLists, expectedValid...)
	pageLists = append(pageLists, expectedInvalid...)

	if sum := P1Solution(rules, pageLists); sum != 143 {
		t.Errorf("expected sum to be 143 but was: %d", sum)
	}
}

func TestIsValid(t *testing.T) {
	for _, in := range expectedValid {
		if v, _, _ := isValid(rules, in); !v {
			t.Errorf("expected result to be valid: %v", in)
		}
	}

	for _, in := range expectedInvalid {
		if v, _, _ := isValid(rules, in); v {
			t.Errorf("expected result to be invalid: %v", in)
		}
	}
}

func TestParseInput(t *testing.T) {
	input := common.ReadFile("./input-test.txt")
	actualRules, _ := ParseInput(input)
	if !reflect.DeepEqual(rules, actualRules) {
		t.Error("expected the rules to be equal")
	}
}

func TestP2Solution(t *testing.T) {
	pageLists := make([][]int, 0)
	pageLists = append(pageLists, expectedValid...)
	pageLists = append(pageLists, expectedInvalid...)

	if sum := P2Solution(rules, pageLists); sum != 123 {
		t.Errorf("expected sum to be 123 but was: %d", sum)
	}
}
