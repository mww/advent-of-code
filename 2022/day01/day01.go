package day01

import (
	"fmt"
	"sort"

	"theelements.org/advent-of-code/common"
)

func P1() {
	totals := getTotalsPerElf()

	fmt.Printf("The most calories is: %d\n", totals[0])
}

func P2() {
	totals := getTotalsPerElf()

	sum := totals[0] + totals[1] + totals[2]
	fmt.Printf("Top 3 are: %d, %d, %d\n", totals[0], totals[1], totals[2])
	fmt.Printf("sum: %d\n", sum)
}

func getTotalsPerElf() []int {
	lines := common.ReadFile("./2022/day01/part-1-input.txt")
	//lines := common.ReadFile("./dec01/test-input.txt")
	// Add a blank line to the end of the input so that we trigger summing
	// all of the calories for the final elf.
	lines = append(lines, "")

	res := make([]int, 0, 100)

	elf := 1
	current := 0
	for _, l := range lines {
		if l == "" {
			fmt.Printf("elf %d is carrying %d calories\n", elf, current)
			res = append(res, current)

			elf++
			current = 0
		} else {
			current += common.Atoi(l)
		}
	}

	//log.Printf("before sort: %v", res)
	sort.Slice(res, func(i, j int) bool {
		return res[j] < res[i]
	})
	//log.Printf("after sort: %v", res)

	return res
}
