package day05

import (
	"log"
	"strings"

	"theelements.org/advent-of-code/common"
)

func P1() {
	input := common.ReadFile("./2024/day05/input.txt")
	rules, pageLists := ParseInput(input)

	sum := P1Solution(rules, pageLists)
	log.Printf("The sum of the middle pages is: %d", sum)
}

func P1Solution(rules map[int]map[int]any, pageLists [][]int) int {
	sum := 0
	for _, p := range pageLists {
		if v, _, _ := isValid(rules, p); v {
			middle := (len(p) / 2)
			sum += p[middle]
			//log.Printf("%v - middle value: %d", p, p[middle])
		}
	}
	return sum
}

func isValid(rules map[int]map[int]any, pages []int) (bool, int, int) {
	// For each page, look for any rules that exist
	for i, p := range pages {
		ruleSet, ok := rules[p]
		if !ok {
			continue
		}
		// The ruleSet contains all the pages that must be after the current page.
		// Walk through all the pages before the current one and see if any are in
		// the rule set.
		for j := i - 1; j >= 0; j-- {
			if _, contains := ruleSet[pages[j]]; contains {
				//log.Printf("order: %v, is not valid because of rule: %d|%d", pages, p, pages[j])
				return false, i, j
			}
		}
	}
	return true, 0, 0
}

func P2() {
	input := common.ReadFile("./2024/day05/input.txt")
	rules, pageLists := ParseInput(input)
	sum := P2Solution(rules, pageLists)
	log.Printf("the sum after reordering is: %d", sum)
}

func P2Solution(rules map[int]map[int]any, pageLists [][]int) int {
	sum := 0
	for _, p := range pageLists {
		work := make([]int, len(p))
		copy(work, p)

		if v, i, j := isValid(rules, work); !v {
			//log.Printf("start page list: %v", p)

			for !v {
				temp := work[i]
				work[i] = work[j]
				work[j] = temp
				//log.Printf("swapped: %d and %d - %v", i, j, work)
				v, i, j = isValid(rules, work)
			}

			//log.Printf("final page list: %v", work)
			middle := (len(work) / 2)
			sum += work[middle]
			//log.Printf("%v - middle value: %d", work, work[middle])
		}
	}
	return sum
}

func ParseInput(input []string) (map[int]map[int]any, [][]int) {
	rules := make(map[int]map[int]any)
	pageLists := make([][]int, 0)

	inRules := true
	for _, i := range input {
		if i == "" {
			inRules = false
			continue
		}

		if inRules {
			parts := strings.Split(i, "|")
			if len(parts) != 2 {
				log.Fatalf("bad rule format: '%s'", i)
			}
			page := common.Atoi(parts[0])
			before := common.Atoi(parts[1])

			p, ok := rules[page]
			if !ok {
				p = make(map[int]any)
				rules[page] = p
			}
			p[before] = true
		} else {
			parts := strings.Split(i, ",")
			l := make([]int, len(parts))
			for x, v := range parts {
				l[x] = common.Atoi(v)
			}
			pageLists = append(pageLists, l)
		}
	}

	return rules, pageLists
}
