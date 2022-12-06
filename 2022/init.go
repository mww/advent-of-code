package year2022

import (
	"theelements.org/advent-of-code/2022/dec01"
	"theelements.org/advent-of-code/2022/dec02"
	"theelements.org/advent-of-code/2022/dec03"
	"theelements.org/advent-of-code/2022/dec04"
	"theelements.org/advent-of-code/common"
)

func init() {
	common.RegisterSolution(&common.ID{Year: 2022, Day: 1, Part: 1}, dec01.P1)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 1, Part: 2}, dec01.P2)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 2, Part: 1}, dec02.P1)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 2, Part: 2}, dec02.P2)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 3, Part: 1}, dec03.P1)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 3, Part: 2}, dec03.P2)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 4, Part: 1}, dec04.P1)
	common.RegisterSolution(&common.ID{Year: 2022, Day: 4, Part: 2}, dec04.P2)
}
