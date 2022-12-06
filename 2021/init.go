package year2021

import (
	"theelements.org/advent_of_code/2021/day06"
	"theelements.org/advent_of_code/common"
)

func init() {
	common.RegisterSolution(&common.ID{Year: 2021, Day: 6, Part: 1}, day06.P1)
	common.RegisterSolution(&common.ID{Year: 2021, Day: 6, Part: 2}, day06.P2)
}
