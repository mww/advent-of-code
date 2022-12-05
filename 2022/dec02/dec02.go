package dec02

import (
	"fmt"
	"log"
	"strings"

	"theelements.org/advent_of_code/common"
)

type RPS int

const (
	Rock     RPS = 1
	Paper        = 2
	Scissors     = 3

	win  int = 6
	lose int = 0
	draw int = 3
)

// r is your value, o is your opponents, the result is the total score of the round
func (r RPS) score(o RPS) int {
	value := 1000
	if r == Rock {
		if o == Rock {
			value = draw
		} else if o == Paper {
			value = lose
		} else if o == Scissors {
			value = win
		} else {
			log.Fatalf("unknown o value: %v", o)
		}
	} else if r == Paper {
		if o == Rock {
			value = win
		} else if o == Paper {
			value = draw
		} else if o == Scissors {
			value = lose
		} else {
			log.Fatalf("unknown o value: %v", o)
		}
	} else if r == Scissors {
		if o == Rock {
			value = lose
		} else if o == Paper {
			value = win
		} else if o == Scissors {
			value = draw
		} else {
			log.Fatalf("unknown o value: %v", o)
		}
	} else {
		log.Fatalf("unknown value: %v", r)
	}

	return int(r) + value
}

func getRPS(i string) RPS {
	if i == "A" || i == "X" {
		return Rock
	} else if i == "B" || i == "Y" {
		return Paper
	} else if i == "C" || i == "Z" {
		return Scissors
	} else {
		log.Fatalf("unknown input: %v", i)
		return Rock
	}
}

func P1() {
	lines := common.ReadFile("./2022/dec02/part-1-input.txt")

	score := 0
	for _, l := range lines {
		p := strings.Split(l, " ")
		if len(p) != 2 {
			log.Fatalf("wrong number of parts when splitting string: '%v'", l)
		}

		o := getRPS(p[0]) // opponent
		m := getRPS(p[1]) // me

		score += m.score(o)
	}

	fmt.Printf("total score: %d\n", score)
}

func P2() {
	lines := common.ReadFile("./2022/dec02/part-1-input.txt")

	score := 0
	for _, l := range lines {
		p := strings.Split(l, " ")
		if len(p) != 2 {
			log.Fatalf("wrong number of parts when splitting string: '%v'", l)
		}

		o := getRPS(p[0]) // opponent

		// decode the desired result
		var m RPS
		switch p[1] {
		case "X":
			// lose
			switch o {
			case Rock:
				m = Scissors
			case Paper:
				m = Rock
			case Scissors:
				m = Paper
			default:
				log.Fatalf("unknown o type: %v", o)
			}
		case "Y":
			// draw
			switch o {
			case Rock:
				m = Rock
			case Paper:
				m = Paper
			case Scissors:
				m = Scissors
			default:
				log.Fatalf("unknown o type: %v", o)
			}
		case "Z":
			// win
			switch o {
			case Rock:
				m = Paper
			case Paper:
				m = Scissors
			case Scissors:
				m = Rock
			default:
				log.Fatalf("unknown o type: %v", o)
			}
		}

		score += m.score(o)
	}
	fmt.Printf("total score: %d\n", score)
}
