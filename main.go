package main

import (
	"flag"
	"log"
	"time"

	_ "theelements.org/advent-of-code/2021"
	_ "theelements.org/advent-of-code/2022"
	_ "theelements.org/advent-of-code/2024"
	"theelements.org/advent-of-code/common"
)

func main() {
	year := flag.Int("year", 0, "The year to run")
	day := flag.Int("day", 0, "The day of code to run")
	part := flag.Int("part", 0, "The part of the code to run, 1 or 2")
	sample := flag.Bool("sample", false, "Run the sample version instead of the real version")

	flag.Parse()

	if !common.IsYearSupported(*year) {
		log.Fatalf("Year %d is not supported", *year)
	}

	if *day == 0 {
		log.Fatal("You must specify the day to run")
	}

	if *part != 1 && *part != 2 {
		log.Fatal("You must specify the part to run")
	}

	solution := common.GetSolution(*year, *day, *part, *sample)

	start := time.Now()
	defer func() {
		log.Printf("took: %v", time.Since(start))
	}()
	solution()
}
