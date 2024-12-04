package day03

import (
	"log"
	"regexp"
	"strings"

	"theelements.org/advent-of-code/common"
)

var re = regexp.MustCompile(`mul\((?P<X>\d{1,3}),(?P<Y>\d{1,3})\)`)

func P1() {
	input := common.ReadFile("./2024/day03/input.txt")
	sum := P1Solution(input)
	log.Printf("The sum is: %d", sum)
}

func P1Solution(input []string) int64 {
	sum := int64(0)
	for _, l := range input {
		matches := re.FindAllStringSubmatch(l, -1)
		for _, m := range matches {
			x := m[re.SubexpIndex("X")]
			y := m[re.SubexpIndex("Y")]

			xn := common.Atoi(x)
			yn := common.Atoi(y)

			sum = sum + int64(xn*yn)
		}
	}
	return sum
}

func P2() {
	input := common.ReadFile("./2024/day03/input.txt")
	sum := P2Solution(input)
	log.Printf("The sum is: %d", sum)
}

func P2Solution(input []string) int64 {
	enabled := true
	parts := make([]string, 0)
	for _, i := range input {
		for {
			//log.Printf("i: %s", i)
			if enabled {
				ndx := strings.Index(i, "don't()")
				if ndx != -1 {
					before := i[:ndx]
					parts = append(parts, before)
					i = i[ndx:]
					//log.Printf("before: %s", before)
					//log.Printf("after: %s", i)
					enabled = false
				} else {
					parts = append(parts, i)
					break
				}
			} else {
				ndx := strings.Index(i, "do()")
				if ndx != -1 {
					i = i[ndx+4:]
					enabled = true
					//log.Printf("after: %s", i)
				} else {
					break
				}
			}
		}
	}

	sum := int64(0)
	for _, l := range parts {
		matches := re.FindAllStringSubmatch(l, -1)
		for _, m := range matches {
			x := m[re.SubexpIndex("X")]
			y := m[re.SubexpIndex("Y")]

			xn := common.Atoi(x)
			yn := common.Atoi(y)

			sum = sum + int64(xn*yn)
		}
	}
	return sum
}
