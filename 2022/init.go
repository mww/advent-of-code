package year2022

import (
	"theelements.org/advent_of_code/2022/dec01"
	"theelements.org/advent_of_code/common"
)

func init() {
	common.RegisterSolution(&common.ID{Year: 2022, Day: 1, Part: 1}, dec01.P1)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 1, Part: 2}, dec01.P2)
}
