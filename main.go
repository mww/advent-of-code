package main

import (
	"flag"
	"log"
	"time"

	_ "theelements.org/advent_of_code/2022"
	"theelements.org/advent_of_code/common"
)

func main() {
	supportedYears := map[int]bool{
		2022: true,
	}

	year := flag.Int("year", 0, "The year to run")
	day := flag.Int("day", 0, "The day of code to run")
	part := flag.Int("part", 0, "The part of the code to run, 1 or 2")

	flag.Parse()

	if _, ok := supportedYears[*year]; !ok {
		log.Fatalf("Year %d is not supported", *year)
	}

	if *day == 0 {
		log.Fatal("You must specify the day to run")
	}

	if *part != 1 && *part != 2 {
		log.Fatal("You must specify the part to run")
	}

	sol := common.GetSolution(*year, *day, *part)

	s := time.Now()
	defer func() {
		log.Printf("took: %v", time.Since(s))
	}()
	sol()
}
