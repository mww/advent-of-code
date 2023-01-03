package day07

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"theelements.org/advent-of-code/common"
)

func P1() {
	data := readData("./2021/day07/input.txt")
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	//log.Printf("sorted data: %v", data)

	// Find the median value. Use an array in case there are an even
	// number of elements and we need to try different medians
	med := make([]int, 0, 2)
	if len(data)%2 == 0 {
		// even number of elements
		mid := len(data) / 2
		med = append(med, data[mid-1], data[mid])
	} else {
		// odd number of elements
		mid := len(data) / 2
		med = append(med, data[mid])
	}

	// calculate the fuel cost for moving each element to the median
	pos := 0
	cost := math.MaxInt
	for _, m := range med {
		//log.Printf("using m: %d", m)
		c := 0
		for _, d := range data {
			c += int(math.Abs(float64(d - m)))
			if c > cost {
				break
			}
		}

		if c < cost {
			cost = c
			pos = m
		}
	}

	fmt.Printf("pos: %d, cost: %d\n", pos, cost)
}

func P2() {
	data := readData("./2021/day07/input.txt")

	max := 0
	candidates := make([]int, 0, 2)
	// Calculate the average value
	sum := 0
	for _, d := range data {
		sum += d
		if d > max {
			max = d
		}
	}
	avg := float64(sum) / float64(len(data))
	candidates = append(candidates, int(math.Ceil(avg)))
	candidates = append(candidates, int(math.Floor(avg)))
	fmt.Printf("candidates: %v\n", candidates)

	// Make a table of movement costs
	costTable := make([]int, max)
	for i := 1; i < max; i++ {
		costTable[i] = costTable[i-1] + i
	}

	// calculate the fuel cost for moving each element to the median
	pos := 0
	cost := math.MaxInt
	for _, m := range candidates {
		c := 0
		for _, d := range data {
			i := int(math.Abs(float64(d - m)))
			c += costTable[i]
			if c > cost {
				break
			}
		}

		if c < cost {
			cost = c
			pos = m
		}
	}

	fmt.Printf("pos: %d, cost:%d\n", pos, cost)
}

func readData(f string) []int {
	res := make([]int, 0, 1000)

	lines := common.ReadFile(f)
	for _, l := range lines {
		parts := strings.Split(l, ",")
		for _, p := range parts {
			i := common.Atoi(p)
			res = append(res, i)
		}
	}

	return res
}
