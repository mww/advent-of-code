package day03

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"theelements.org/advent-of-code/common"
)

const dataLength = 12

func P1() {
	digitCountZeros := make([]int, dataLength)
	digitCountOnes := make([]int, dataLength)

	lines := common.ReadFile("./2021/day03/input.txt")
	for _, l := range lines {
		for i, r := range l {
			if r == '0' {
				digitCountZeros[i]++
			} else if r == '1' {
				digitCountOnes[i]++
			} else {
				log.Fatalf("unexpected input line: '%s'", l)
			}
		}
	}

	var gamma strings.Builder
	for i := 0; i < dataLength; i++ {
		if digitCountZeros[i] > digitCountOnes[i] {
			gamma.WriteRune('0')
		} else {
			gamma.WriteRune('1')
		}
	}

	gammaStr := gamma.String()
	//log.Printf("gamma: %s", gammaStr)

	g, err := strconv.ParseInt(gammaStr, 2, 64)
	if err != nil {
		log.Fatalf("error parsing gamma value: %v", err)
		return
	}

	// This only works if the data size is 12, should have a better solution
	e := g ^ 0b111111111111

	fmt.Printf("gamma: %d, epsilon: %d, product: %d\n", g, e, g*e)
}

func P2() {
	lines := common.ReadFile("./2021/day03/input.txt")

	o2 := doFilter("o2", lines, func(zeros, ones int, c rune) bool {
		if ones >= zeros && c == '1' {
			return true
		} else if zeros > ones && c == '0' {
			return true
		}
		return false
	})

	co2 := doFilter("co2", lines, func(zeros, ones int, c rune) bool {
		if ones < zeros && c == '1' {
			return true
		} else if zeros <= ones && c == '0' {
			return true
		}
		return false
	})

	log.Printf("o2 value: %d, co2 value: %d, product: %d", o2, co2, o2*co2)
}

type accept func(zeros, ones int, c rune) bool

func doFilter(name string, data []string, acc accept) int64 {
	for pos := 0; pos < dataLength; pos++ {
		data = filter(data, pos, acc)
		if len(data) == 1 {
			break
		}
	}

	if len(data) != 1 {
		log.Fatalf("did not filter %s down to a single value", name)
	}

	v, err := strconv.ParseInt(data[0], 2, 32)
	if err != nil {
		log.Fatalf("error parsing %s number: %v - '%s'", name, err, data[0])
	}
	return v
}

func filter(data []string, pos int, acc accept) []string {
	// Iterate through all the data and count the 1s and 0s.
	zeros := 0
	ones := 0
	for _, d := range data {
		r := []rune(d)[pos]
		if r == '0' {
			zeros++
		} else if r == '1' {
			ones++
		} else {
			log.Fatalf("expected rune in input: %s", d)
		}
	}

	// starting size of 100 just to reduce the total number of allocations required
	res := make([]string, 0, 100)

	// Now look through the data and decide which entries should be included
	for _, d := range data {
		c := []rune(d)[pos]
		if acc(zeros, ones, c) {
			res = append(res, d)
		}
	}

	//log.Printf("%d items left after iteration %d", len(res), pos)
	return res
}
