package day05

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"theelements.org/advent-of-code/common"
)

func P1() {
	stacks, commands := parseInput("./2022/day05/input.txt")

	for _, c := range commands {
		//log.Printf("executing command: move %d from %d to %d", c.num, c.from, c.to)
		for i := 0; i < c.num; i++ {
			// the commands are 1 based, the arrays are 0 based, so subtract one
			// from the number to get the stack index
			t := stacks[c.from-1].pop()
			stacks[c.to-1].push(t)
		}
	}

	for i, s := range stacks {
		log.Printf("stack %d: %v", i+1, &s)
	}

	fmt.Print("top elements: ")
	for _, s := range stacks {
		fmt.Printf("%c", s.peek())
	}
	fmt.Println()
}

func P2() {
	stacks, commands := parseInput("./2022/day05/input.txt")

	for _, c := range commands {
		// the commands are 1 based, the arrays are 0 based, so subtract one
		// from the number to get the stack index
		//log.Printf("executing command: move %d from %d to %d", c.num, c.from, c.to)
		popped := stacks[c.from-1].popN(c.num)
		for _, i := range popped {
			stacks[c.to-1].push(i)
		}
	}

	for i, s := range stacks {
		log.Printf("stack %d: %v", i+1, &s)
	}

	fmt.Print("top elements: ")
	for _, s := range stacks {
		fmt.Printf("%c", s.peek())
	}
	fmt.Println()
}

type stack []rune

func (s *stack) String() string {
	x := make([]string, 0, len(*s))
	for _, r := range *s {
		x = append(x, fmt.Sprintf("[%c]", r))
	}

	return strings.Join(x, " ")
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s *stack) pop() rune {
	if s.isEmpty() {
		log.Fatalf("trying to pop from an empty stack")
	}

	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e
}

func (s *stack) popN(n int) []rune {
	if s.isEmpty() {
		log.Fatalf("trying to popN from an empty stack")
	}

	end := len(*s)
	start := end - n

	r := (*s)[start:end]
	*s = (*s)[:start]
	return r
}

func (s *stack) peek() rune {
	if s.isEmpty() {
		log.Fatalf("trying to peek on an empty stack")
	}
	i := len(*s) - 1
	return (*s)[i]
}

type command struct {
	num  int
	from int
	to   int
}

func parseInput(f string) ([]stack, []command) {
	lines := common.ReadFile(f)

	// The first several lines are all stack states.
	// Each line is exactly the same length, 3 characters for each crate ([A])
	// and a space before the next crate. The final crate does not have a trailing space.
	// So if the first line has a length of 11 we know there are 3 crates.
	// We can add 1 to length and divide by 4 to to get the number of crates.
	numCrates := (len(lines[0]) + 1) / 4
	//log.Printf("num crates: %d", numCrates)

	temp := make([]stack, numCrates)
	stacks := make([]stack, numCrates)

	var i int
	for i = 0; i < len(lines); i++ {
		l := lines[i]

		if l == "" {
			// This is the blank line that separates the stack state
			// from the commands.
			break
		}

		open := false
		// Look at each character one at a time
		for x, r := range l {
			if r == ' ' {
				continue
			}

			if r == '[' {
				open = true
				continue
			}

			if r == ']' {
				open = false
				continue
			}

			if open {
				num := x / 4
				temp[num].push(r)
			} else {
				// This should be the line that numbers the stacks
				//log.Printf("line: '%s'", l)
				break
			}
		}

	}

	// The stacks are built from the top down, so now we need to reverse them
	for i, s := range temp {
		for !s.isEmpty() {
			t := s.pop()
			stacks[i].push(t)
		}
		//log.Printf("stack %d: %v", i+1, &stacks[i])
	}

	// now parse the commands
	commands := make([]command, 0, 100)
	re := regexp.MustCompile(`^move (?P<num>\d+) from (?P<from>\d+) to (?P<to>\d+)$`)
	// pick up where we left off in the input
	for i = i + 1; i < len(lines); i++ {
		l := lines[i]
		results := re.FindStringSubmatch(l)
		if results == nil || len(results) != 4 {
			log.Fatalf("error parsing command: '%s'", l)
		}

		c := command{
			num:  common.Atoi(results[1]),
			from: common.Atoi(results[2]),
			to:   common.Atoi(results[3]),
		}
		commands = append(commands, c)
	}

	return stacks, commands
}
