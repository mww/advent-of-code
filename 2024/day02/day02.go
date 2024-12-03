package day02

import (
	"log"
	"strings"

	"theelements.org/advent-of-code/common"
)

func P1Sample() {
	safeReports := P1Solution("./2024/day02/input-sample.txt")
	if safeReports != 2 {
		log.Printf("wrong answer, expected 2, got: %d", safeReports)
	} else {
		log.Printf("correct answer")
	}
}

func P1() {
	sum := P1Solution("./2024/day02/input.txt")
	log.Printf("The number of safe reports is: %d", sum)
}

func P1Solution(input string) int {
	data := readInput(input)
	safe := 0

	for _, line := range data {
		if isLineSafe(line) {
			safe++
		}
	}

	return safe
}

func isLineSafe(line []int) bool {
	inc := true
	if line[0] > line[1] {
		inc = false
	}

	for i := 1; i < len(line); i++ {
		if inc {
			if line[i] < line[i-1] {
				return false
			}
		} else {
			if line[i] > line[i-1] {
				return false
			}
		}

		step := line[i] - line[i-1]
		if step < 0 {
			step = step * -1
		}
		if step == 0 {
			return false
		}
		if step > 3 {
			return false
		}
	}

	return true
}

func P2Sample() {
	safeReports := P2Solution("./2024/day02/input-sample.txt")
	if safeReports != 4 {
		log.Printf("wrong answer, expected 4, got: %d", safeReports)
	} else {
		log.Printf("correct answer")
	}
}

func P2() {
	sum := P2Solution("./2024/day02/input.txt")
	log.Printf("The number of safe reports is: %d", sum)
}

func P2Solution(input string) int {
	data := readInput(input)
	safe := 0

	for _, line := range data {
		if part2IsLineSafe(line) {
			safe++
		}
	}

	return safe
}

func part2IsLineSafe(line []int) bool {
	//log.Printf("start line: %v", line)
	safe := isLineSafe(line)
	if !safe {
		for i := 0; i < len(line); i++ {
			lineCopy := make([]int, len(line))
			copy(lineCopy, line)

			//log.Printf("line: %v, i: %d", lineCopy, i)
			lineCopy = append(lineCopy[:i], lineCopy[i+1:]...)
			//log.Printf("lineCopy after removal: %v", lineCopy)
			if isLineSafe(lineCopy) {
				//log.Printf("safe after removal: %v", lineCopy)
				safe = true
				break
			}
		}
	}
	return safe
}

func readInput(file string) [][]int {
	lines := common.ReadFile(file)
	result := make([][]int, 0, len(lines))

	for _, l := range lines {
		parts := strings.Split(l, " ")
		data := make([]int, len(parts))
		for i, v := range parts {
			data[i] = common.Atoi(v)
		}
		result = append(result, data)
	}

	return result
}
