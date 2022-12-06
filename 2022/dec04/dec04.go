package dec04

import (
	"fmt"
	"log"
	"strings"

	"theelements.org/advent-of-code/common"
)

func P1() {
	lines := common.ReadFile("./2022/dec04/input.txt")

	count := 0
	for _, l := range lines {
		parts := strings.Split(l, ",")
		if len(parts) != 2 {
			log.Fatalf("wrong number of parts in line '%s'", l)
		}

		a := getSet(parts[0])
		b := getSet(parts[1])

		if len(a) > len(b) {
			if isSubset(b, a) {
				count++
			}
		} else {
			if isSubset(a, b) {
				count++
			}
		}
	}

	fmt.Printf("count: %d\n", count)
}

func P2() {
	lines := common.ReadFile("./2022/dec04/input.txt")

	overlap := 0
	for _, l := range lines {
		parts := strings.Split(l, ",")
		if len(parts) != 2 {
			log.Fatalf("wrong number of parts in line '%s'", l)
		}

		a := getSet(parts[0])
		b := getSet(parts[1])

		if hasOverlap(a, b) {
			overlap++
		}
	}

	fmt.Printf("has overlap: %d\n", overlap)
}

// parses a string like "2-4" into a set of ints containing {2, 3, 4}
func getSet(s string) map[int]any {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		log.Fatalf("wrong format for set '%s'", s)
	}

	start := common.Atoi(parts[0])
	end := common.Atoi(parts[1])

	res := make(map[int]any)
	for i := start; i <= end; i++ {
		res[i] = true
	}

	return res
}

// returns true if a is a subset of b
func isSubset(a, b map[int]any) bool {
	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	return true
}

// returns true if a and b have any elements in common
func hasOverlap(a, b map[int]any) bool {
	for k := range a {
		if _, ok := b[k]; ok {
			return true
		}
	}
	return false
}
