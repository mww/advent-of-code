package day07

import (
	"fmt"
	"log"
	"strings"

	"theelements.org/advent-of-code/common"
)

func P1() {
	input := common.ReadFile("./2024/day07/input.txt")
	s := P1Solution(input)
	log.Printf("The sum of the possible lines are: %d", s)
}

func P1Solution(input []string) int64 {
	possOps := []rune{'*', '+'}

	sum := int64(0)
	for _, line := range input {
		parts := strings.Split(line, " ")
		p0 := strings.TrimSuffix(parts[0], ":")
		total := common.Atoi64(p0)
		vals := make([]int64, len(parts)-1)
		for i := 1; i < len(parts); i++ {
			vals[i-1] = common.Atoi64(parts[i])
		}

		valid, _ := couldBeTrue(total, vals, possOps)
		if valid {
			sum += total
		}
	}
	return sum
}

func P2() {
	input := common.ReadFile("./2024/day07/input.txt")
	s := P2Solution(input)
	log.Printf("The sum of the possible lines are: %d", s)
}

func P2Solution(input []string) int64 {
	// since we are using runes, use '|' instead of '||'
	possOps := []rune{'*', '+', '|'}

	sum := int64(0)
	for _, line := range input {
		parts := strings.Split(line, " ")
		p0 := strings.TrimSuffix(parts[0], ":")
		total := common.Atoi64(p0)
		vals := make([]int64, len(parts)-1)
		for i := 1; i < len(parts); i++ {
			vals[i-1] = common.Atoi64(parts[i])
		}

		valid, _ := couldBeTrue(total, vals, possOps)
		if valid {
			sum += total
		}
	}
	return sum
}

// Returns if the input could be true, and if true then it returns
// a list of operators that makes it true.
func couldBeTrue(sum int64, values []int64, possOps []rune) (bool, []rune) {
	acc := make([][]rune, 0)
	cur := make([]rune, 0)
	acc = generateOpsOptions(acc, cur, len(values)-1, possOps)

	for _, ops := range acc {
		v := values[0]
		for i, op := range ops {
			if op == '*' {
				v *= values[i+1]
			} else if op == '+' {
				v += values[i+1]
			} else if op == '|' {
				s := fmt.Sprintf("%d%d", v, values[i+1])
				v = common.Atoi64(s)
			} else {
				log.Fatalf("unsupported operation: %c", op)
			}
		}
		if v == sum {
			return true, ops
		}
	}
	return false, nil
}

func generateOpsOptions(acc [][]rune, cur []rune, left int, possOps []rune) [][]rune {
	for _, p := range possOps {
		c := make([]rune, len(cur))
		copy(c, cur)
		c = append(c, p)
		if left == 1 {
			acc = append(acc, c)
		} else {
			acc = generateOpsOptions(acc, c, left-1, possOps)
		}
	}
	return acc
}
