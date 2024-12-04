package day03

import "testing"

func TestP1Solution(t *testing.T) {
	input := []string{
		"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
	}

	result := P1Solution(input)
	if result != 161 {
		t.Errorf("expected result to be 161 but was %d", result)
	}
}

func TestP2Solution(t *testing.T) {
	input := []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}

	result := P2Solution(input)
	if result != 48 {
		t.Errorf("expected result to be 48 but was %d", result)
	}
}
