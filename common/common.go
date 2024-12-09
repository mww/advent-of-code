package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ID struct {
	Year   int
	Day    int
	Part   int
	Sample bool
}

type Solution func()

var solutions map[string]Solution
var supportedYears map[int]bool

func init() {
	solutions = make(map[string]Solution)
	supportedYears = make(map[int]bool)
}

func RegisterSolution(id *ID, sol Solution) {
	key := fmt.Sprintf("%d-%02d-%d-%v", id.Year, id.Day, id.Part, id.Sample)
	solutions[key] = sol
	supportedYears[id.Year] = true
}

func GetSolution(year, day, part int, sample bool) Solution {
	key := fmt.Sprintf("%d-%02d-%d-%v", year, day, part, sample)
	sol, ok := solutions[key]
	if !ok {
		log.Fatalf("no solution registered for year %d, day %02d, part %d", year, day, part)
	}
	return sol
}

func IsYearSupported(year int) bool {
	_, ok := supportedYears[year]
	return ok
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

func Atoi64(i string) int64 {
	v, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
