package day06

import (
	"fmt"
	"strings"

	"theelements.org/advent_of_code/common"
)

func P1() {
	const days = 80 // part 1
	data := readData("./2021/day06/input.txt")

	//log.Printf("initial input is %d items", len(data))
	for i := 0; i < days; i++ {
		//log.Printf("day: %2d", i+1)
		add := 0
		for x := range data {
			if data[x] == 0 {
				data[x] = 6
				add++
			} else {
				data[x]--
			}
		}

		for x := 0; x < add; x++ {
			data = append(data, 8)
		}

		//log.Printf("After %2d days: %v", i+1, data)
	}

	fmt.Printf("total fish: %d\n", len(data))
}

func P2() {
	const days = 256
	data := readData("./2021/day06/input.txt")

	// The solution in part 1 won't scale - the array gets too big.
	//
	// build a table that tells us how many fish are added to the population
	// for any fish born on a given day.
	// The general formula is 1 + f(d+9) + f(d+9+7) + f(d+9+7+7) all the way
	// until d+9+7+7... is greater than our max day.
	// We build the table from the end so that we can always reference an earlier day

	// days+1 because we want the table from 0-days
	table := make([]int64, days+1)

	for i := days; i >= 0; i-- {
		//log.Printf("day: %d", i)
		sum := int64(1) // For this fish that was just born

		j := i + 9
		for j <= days {
			//log.Printf("\tadding fish for day: %d -> %d", j, table[j])
			sum += table[j]
			j += 7
		}

		//log.Printf("\ttotal fish for day %d: %d", i, sum)
		table[i] = sum
	}

	//log.Printf("%v\n", table)

	sum := int64(0)
	for _, d := range data {
		//log.Printf("data: %d", d)
		sum++ // Add one for this fish

		// Now figure out when the next fish will be born.
		j := d + 1
		for j <= days {
			//log.Printf("adding fish for day: %d -> %d", j, table[j])
			sum += table[j]
			j += 7
		}
	}

	fmt.Printf("After %d days: %d fish\n", days, sum)
}

func readData(f string) []int {
	data := make([]int, 0, 100)

	lines := common.ReadFile(f)
	for _, l := range lines {
		parts := strings.Split(l, ",")
		for _, p := range parts {
			i := common.Atoi(p)
			data = append(data, i)
		}
	}

	return data
}
