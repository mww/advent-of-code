package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Solution func()

var solutions map[string]Solution

func init() {
	solutions = make(map[string]Solution)
}

func RegisterSolution(year, day, part int, sol Solution) {
	key := fmt.Sprintf("%d-%02d-%d", year, day, part)
	solutions[key] = sol
}

func GetSolution(year, day, part int) Solution {
	key := fmt.Sprintf("%d-%02d-%d", year, day, part)
	sol, ok := solutions[key]
	if !ok {
		log.Fatalf("no solution registered for year %d, day %02d, part %d", year, day, part)
	}
	return sol
}

func ReadFile(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}

func Atoi(num string) int {
	i, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
