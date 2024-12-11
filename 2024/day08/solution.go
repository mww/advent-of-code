package day08

import (
	"fmt"
	"log"

	"theelements.org/advent-of-code/common"
)

func P1() {
	input := common.ReadFile("./2024/day08/input.txt")
	data := readData(input)
	h := len(input)
	w := len(input[0])

	n := P1Solution(w, h, data)
	log.Printf("The number of antinodes is: %d", n)
}

// w and h are the max width and heights so we can determine
// if an antinode is outside of the map
func P1Solution(w, h int, data map[rune][]loc) int {
	antinodes := make(map[string]any)

	for _, locs := range data {
		for i := 0; i < len(locs); i++ {
			for j := 0; j < len(locs); j++ {
				if i == j {
					continue
				}

				dr := locs[j].row - locs[i].row
				dc := locs[j].col - locs[i].col

				ar := locs[i].row - dr
				ac := locs[i].col - dc

				if ar >= 0 && ar < h && ac >= 0 && ac < w {
					antinodes[fmt.Sprintf("(%d, %d)", ac, ar)] = true
				}
			}
		}
	}

	return len(antinodes)
}

func P2() {
	input := common.ReadFile("./2024/day08/input.txt")
	data := readData(input)
	h := len(input)
	w := len(input[0])

	n := P2Solution(w, h, data)
	log.Printf("The number of antinodes is: %d", n)
}

func P2Solution(w, h int, data map[rune][]loc) int {
	antinodes := make(map[string]any)

	for _, locs := range data {
		for i := 0; i < len(locs); i++ {
			for j := 0; j < len(locs); j++ {
				if i == j {
					continue
				}
				antinodes[fmt.Sprintf("(%d, %d)", locs[i].col, locs[i].row)] = true

				dr := locs[j].row - locs[i].row
				dc := locs[j].col - locs[i].col

				r, c := locs[i].row, locs[i].col
				for {
					r -= dr
					c -= dc

					if r >= 0 && r < h && c >= 0 && c < w {
						antinodes[fmt.Sprintf("(%d, %d)", c, r)] = true
					} else {
						break
					}
				}
			}
		}
	}

	return len(antinodes)
}

func readData(input []string) map[rune][]loc {
	res := make(map[rune][]loc)
	for row, line := range input {
		for col, r := range line {
			if r == '.' {
				continue
			}

			if _, ok := res[r]; !ok {
				locs := make([]loc, 0, 4)
				res[r] = locs
			}

			locs := res[r]
			res[r] = append(locs, loc{row: row, col: col})
		}
	}

	return res
}

type loc struct {
	row, col int
}
