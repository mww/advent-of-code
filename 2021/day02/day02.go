package day02

import (
	"fmt"
	"log"
	"strings"

	"theelements.org/advent-of-code/common"
)

type command struct {
	Dir  string
	Step int
}

func P1() {
	commands := getCommands("./2021/day02/input.txt")

	hPos, depth := 0, 0

	for _, c := range commands {
		if c.Dir == "forward" {
			hPos += c.Step
		} else if c.Dir == "up" {
			depth -= c.Step
		} else if c.Dir == "down" {
			depth += c.Step
		} else {
			log.Fatalf("unsupported direction")
			return
		}
	}

	fmt.Printf("horizontal position: %d, depth: %d, product: %d\n", hPos, depth, hPos*depth)
}

func P2() {
	commands := getCommands("./2021/day02/input.txt")
	hPos, depth, aim := 0, 0, 0

	for _, c := range commands {
		if c.Dir == "forward" {
			hPos += c.Step
			depth += c.Step * aim
		} else if c.Dir == "up" {
			aim -= c.Step
		} else if c.Dir == "down" {
			aim += c.Step
		} else {
			log.Fatalf("unsupported direction")
			return
		}
	}

	log.Printf("horizontal position: %d, depth: %d, product: %d", hPos, depth, hPos*depth)

}

func getCommands(f string) []command {
	commands := make([]command, 0, 1000)
	lines := common.ReadFile(f)
	for _, l := range lines {
		parts := strings.Split(l, " ")
		if len(parts) != 2 {
			log.Fatalf("wrong format for command. got: '%s'", l)
		}

		step := common.Atoi(parts[1])
		commands = append(commands, command{Dir: parts[0], Step: step})
	}

	return commands
}
