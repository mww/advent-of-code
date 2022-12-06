package day01

import (
	"fmt"
	"math"

	"theelements.org/advent-of-code/common"
)

func P1() {
	lines := common.ReadFile("./2021/day01/input.txt")

	prev := math.MaxInt
	count := int(0)
	for _, l := range lines {
		i := common.Atoi(l)
		if i > prev {
			count++
		}
		prev = i
	}

	fmt.Printf("increases: %d\n", count) // Correct answer: 1228
}

// Same as before, but looking at the sum of 3 measurements in a sliding window
// 199  A
// 200  A B
// 208  A B C
// 210    B C D
// 200  E   C D
// 207  E F   D
// 240  E F G
// 269    F G H
// 260      G H
// 263        H
//
// A = 199+200+208 = 607
// B = 200+208+210 = 618
// etc.
func P2() {
	const windowSize = 3

	lines := common.ReadFile("./2021/day01/input.txt")
	data := make([]int, 0, 1000)

	for _, l := range lines {
		i := common.Atoi(l)
		data = append(data, i)
	}

	count := 0
	prev := math.MaxInt
	for i := 0; i <= len(data)-windowSize; i++ {
		cur := data[i] + data[i+1] + data[i+2]
		if cur > prev {
			count++
		}
		prev = cur
	}

	fmt.Printf("increases: %d\n", count) // correct answer: 1257
}
