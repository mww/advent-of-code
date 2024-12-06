package year2024

import (
	"theelements.org/advent-of-code/2024/day01"
	"theelements.org/advent-of-code/2024/day02"
	"theelements.org/advent-of-code/2024/day03"
	"theelements.org/advent-of-code/2024/day04"
	"theelements.org/advent-of-code/2024/day05"
	"theelements.org/advent-of-code/2024/day06"
	"theelements.org/advent-of-code/common"
)

func init() {
	common.RegisterSolution(&common.ID{Year: 2024, Day: 1, Part: 1, Sample: true}, day01.P1Sample)
	common.RegisterSolution(&common.ID{Year: 2024, Day: 1, Part: 1}, day01.P1)
	common.RegisterSolution(&common.ID{Year: 2024, Day: 1, Part: 2, Sample: true}, day01.P2Sample)
	common.RegisterSolution(&common.ID{Year: 2024, Day: 1, Part: 2}, day01.P2)

	common.RegisterSolution(&common.ID{Year: 2024, Day: 2, Part: 1, Sample: true}, day02.P1Sample)
	common.RegisterSolution(&common.ID{Year: 2024, Day: 2, Part: 1}, day02.P1)
	common.RegisterSolution(&common.ID{Year: 2024, Day: 2, Part: 2, Sample: true}, day02.P2Sample)
	common.RegisterSolution(&common.ID{Year: 2024, Day: 2, Part: 2}, day02.P2)

	common.RegisterSolution(&common.ID{Year: 2024, Day: 3, Part: 1}, day03.P1)
	common.RegisterSolution(&common.ID{Year: 2024, Day: 3, Part: 2}, day03.P2)

	common.RegisterSolution(&common.ID{Year: 2024, Day: 4, Part: 1}, day04.P1)
	common.RegisterSolution(&common.ID{Year: 2024, Day: 4, Part: 2}, day04.P2)

	common.RegisterSolution(&common.ID{Year: 2024, Day: 5, Part: 1}, day05.P1)
	common.RegisterSolution(&common.ID{Year: 2024, Day: 5, Part: 2}, day05.P2)

	common.RegisterSolution(&common.ID{Year: 2024, Day: 6, Part: 1}, day06.P1)
	common.RegisterSolution(&common.ID{Year: 2024, Day: 6, Part: 2}, day06.P2)
}
