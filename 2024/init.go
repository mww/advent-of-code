package year2024

import (
	"theelements.org/advent-of-code/2024/day01"
	"theelements.org/advent-of-code/2024/day02"
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
}
