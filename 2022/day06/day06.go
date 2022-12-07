package day06

import (
	"fmt"

	"theelements.org/advent-of-code/common"
)

func P1() {
	lines := common.ReadFile("./2022/day06/input.txt")

	for _, l := range lines {
		findMarker(4, l)
	}
}

func P2() {
	lines := common.ReadFile("./2022/day06/input.txt")

	for _, l := range lines {
		findMarker(14, l)
	}
}

func findMarker(n int, l string) {
	r := []rune(l)

	for i := n - 1; i < len(r); i++ {
		set := make(map[rune]any)

		success := true
		for j := 0; j < n; j++ {
			if _, ok := set[r[i-j]]; ok {
				// We have a duplicate character
				success = false
			} else {
				set[r[i-j]] = true
			}
		}

		if success {
			fmt.Printf("first marker after character %d\n", i+1)
			break
		}
	}
}
