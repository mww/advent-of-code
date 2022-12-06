package day05

import (
	"fmt"
	"log"
	"math"
	"regexp"

	"theelements.org/advent-of-code/common"
)

func P1() {
	lines := parseData("./2021/day05/input.txt")
	maxX, maxY := 0, 0

	for _, l := range lines {
		maxX = max(maxX, l.x1, l.x2)
		maxY = max(maxY, l.y1, l.y2)
	}

	// The arrays are zero based, so add one to the max values
	maxX++
	maxY++
	log.Printf("max X: %d, max Y: %d", maxX, maxY)

	filtered := make([]line, 0, len(lines))
	for _, l := range lines {
		if l.x1 == l.x2 || l.y1 == l.y2 {
			filtered = append(filtered, l)
		}
	}

	log.Printf("number of lines after filtering: %d", len(filtered))
	grid := make([][]int, maxY)
	for i := range grid {
		grid[i] = make([]int, maxX)
	}

	for _, l := range filtered {
		// log.Printf("line: (%d, %d) -> (%d, %d)\n", l.x1, l.y1, l.x2, l.y2)
		if l.x1 == l.x2 {
			// vertical
			y1, y2 := l.y1, l.y2
			// if the points are switched
			if y2 < y1 {
				y1, y2 = l.y2, l.y1
			}
			for y := y1; y <= y2; y++ {
				grid[y][l.x1]++
			}
		} else {
			// horizontal
			x1, x2 := l.x1, l.x2
			// if the points are switched
			if x2 < x1 {
				x1, x2 = l.x2, l.x1
			}
			for x := x1; x <= x2; x++ {
				grid[l.y1][x]++
			}
		}
	}

	overlapping := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 1 {
				overlapping++
			}
		}
	}

	// Print out the grid
	// for i := range grid {
	// 	fmt.Printf("%v\n", grid[i])
	// }

	fmt.Printf("overlapping input: %d\n", overlapping)
}

func P2() {
	lines := parseData("./2021/day05/input.txt")
	maxX, maxY := 0, 0

	for _, l := range lines {
		maxX = max(maxX, l.x1, l.x2)
		maxY = max(maxY, l.y1, l.y2)
	}

	// The arrays are zero based, so add one to the max values
	maxX++
	maxY++
	log.Printf("max X: %d, max Y: %d", maxX, maxY)

	grid := make([][]int, maxY)
	for i := range grid {
		grid[i] = make([]int, maxX)
	}

	for _, l := range lines {
		// log.Printf("line: (%d, %d) -> (%d, %d)\n", l.x1, l.y1, l.x2, l.y2)
		x, y := l.x1, l.y1
		for {
			// log.Printf("marking (%d, %d)", x, y)
			grid[y][x]++

			// See if we are done
			if x == l.x2 && y == l.y2 {
				break
			}

			// Increment/decrement x and y
			if l.x2 > l.x1 {
				x++
			} else if l.x1 > l.x2 {
				x--
			}

			if l.y2 > l.y1 {
				y++
			} else if l.y1 > l.y2 {
				y--
			}
		}
	}

	overlapping := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 1 {
				overlapping++
			}
		}
	}

	// Print out the grid
	// for i := range grid {
	// 	fmt.Printf("%v\n", grid[i])
	// }

	fmt.Printf("overlapping input: %d\n", overlapping)
}

type line struct {
	x1, y1, x2, y2 int
}

func parseData(f string) []line {
	linePattern := regexp.MustCompile(`(?P<x1>\d+),(?P<y1>\d+) -> (?P<x2>\d+),(?P<y2>\d+)`)
	res := make([]line, 0, 100)

	lines := common.ReadFile(f)
	for _, l := range lines {
		results := linePattern.FindStringSubmatch(l)
		if results == nil || len(results) != 5 {
			log.Fatalf("error parsing line: '%s'", l)
		}

		var i line
		i.x1 = common.Atoi(results[1])
		i.y1 = common.Atoi(results[2])
		i.x2 = common.Atoi(results[3])
		i.y2 = common.Atoi(results[4])
		res = append(res, i)
	}

	return res
}

func max(x1, x2, x3 int) int {
	i := math.Max(float64(x2), float64(x3))
	j := math.Max(float64(x1), i)
	return int(j)
}
