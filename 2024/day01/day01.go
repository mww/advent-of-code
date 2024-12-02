package day01

import (
	"log"
	"slices"
	"strings"
	"sync/atomic"

	"theelements.org/advent-of-code/common"
)

func P1Sample() {
	sum := P1Solution("./2024/day01/input-sample.txt")
	if sum != 11 {
		log.Printf("wrong answer, expected 11, got: %d", sum)
	} else {
		log.Printf("correct answer")
	}
	log.Printf("The distance between lists is: %d", sum)
}

func P1() {
	sum := P1Solution("./2024/day01/input.txt")
	log.Printf("The distance between lists is: %d", sum)
}

func P1Solution(input string) int64 {
	l1, l2 := readInput(input)
	slices.SortFunc(l1, sortInts)
	slices.SortFunc(l2, sortInts)

	var sum int64
	for i := 0; i < len(l1); i++ {
		v := l1[i] - l2[i]
		if v < 0 {
			v = v * -1
		}
		sum += int64(v)
	}
	return sum
}

func P2Sample() {
	sum := P2Solution("./2024/day01/input-sample.txt")
	if sum != 31 {
		log.Printf("wrong answer, expected 32, got: %d", sum)
	} else {
		log.Printf("got correct answer")
	}
}

func P2() {
	sum := P2Solution("./2024/day01/input.txt")
	log.Printf("The similarity between lists is: %d", sum)
}

func P2Solution(input string) int64 {
	l1, l2 := readInput(input)

	l2counts := make(map[int]*int64)
	for _, v := range l2 {
		if _, ok := l2counts[v]; !ok {
			var i int64
			l2counts[v] = &i
		}
		atomic.AddInt64(l2counts[v], 1)
	}

	var sum int64
	for _, v := range l1 {
		if _, ok := l2counts[v]; !ok {
			// The value isn't found in list 2, so ignore it
			continue
		}
		sum += (int64(v) * (*l2counts[v]))
	}
	return sum
}

func sortInts(a, b int) int {
	return a - b
}

func readInput(file string) ([]int, []int) {
	lines := common.ReadFile(file)
	l1 := make([]int, len(lines))
	l2 := make([]int, len(lines))

	for i, l := range lines {
		parts := strings.Split(l, "   ")
		if len(parts) != 2 {
			log.Fatalf("input line is not just 2 numbers: '%s'", l)
		}
		l1[i] = common.Atoi(parts[0])
		l2[i] = common.Atoi(parts[1])
	}

	return l1, l2
}
