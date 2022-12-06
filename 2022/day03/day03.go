package day03

import (
	"fmt"

	"theelements.org/advent-of-code/common"
)

func P1() {
	packs := common.ReadFile("./2022/day03/input.txt")

	letterScores := make(map[rune]int)
	for i, l := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		letterScores[l] = i + 1
	}

	score := 0

	for _, p := range packs {
		half := len(p) / 2
		c1 := p[:half]
		c2 := p[half:]

		contents := make(map[rune]bool)
		for _, l := range c1 {
			contents[l] = true
		}

		c2Contents := make(map[rune]bool)
		for _, l := range c2 {
			if _, ok := c2Contents[l]; ok {
				// We've already seen this letter in C2, so ignore it
				continue
			}
			c2Contents[l] = true

			if _, ok := contents[l]; ok {
				// This item is also in c1, so add it's priority to the score
				//log.Printf("duplication item: '%c': %d", l, letterScores[l])
				score += letterScores[l]
			}
		}
	}

	fmt.Printf("sum of priorities: %d\n", score)
}

func P2() {
	packs := common.ReadFile("./2022/day03/input.txt")

	letterScores := make(map[rune]int)
	for i, l := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		letterScores[l] = i + 1
	}

	score := 0

	packSets := make([]map[rune]bool, 0, len(packs))
	for _, p := range packs {
		s := make(map[rune]bool)

		for _, l := range p {
			s[l] = true
		}

		packSets = append(packSets, s)
	}

	// Find the union of the sets, 3 at a time.
	for i := 0; i+3 <= len(packSets); i += 3 {
		p0 := packSets[i]
		p1 := packSets[i+1]
		p2 := packSets[i+2]

		// Look at each item in p0, if it is also in p1 and p2 then it is the
		// badge item. This should always happen, but never more than once.
		for k := range p0 {
			_, ok1 := p1[k]
			_, ok2 := p2[k]
			if ok1 && ok2 {
				// This is the badge item
				//log.Printf("badge item: '%c': %d", k, letterScores[k])
				score += letterScores[k]
				break
			}
		}
	}

	fmt.Printf("sum of priorities: %d\n", score)
}
