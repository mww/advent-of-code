package year2022

import (
	"theelements.org/advent-of-code/2022/day01"
	"theelements.org/advent-of-code/2022/day02"
	"theelements.org/advent-of-code/2022/day03"
	"theelements.org/advent-of-code/2022/day04"
	"theelements.org/advent-of-code/common"
)

func init() {
	common.RegisterSolution(&common.ID{Year: 2022, Day: 1, Part: 1}, day01.P1)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 1, Part: 2}, day01.P2)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 2, Part: 1}, day02.P1)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 2, Part: 2}, day02.P2)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 3, Part: 1}, day03.P1)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 3, Part: 2}, day03.P2)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 4, Part: 1}, day04.P1)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 4, Part: 2}, day04.P2)
}
