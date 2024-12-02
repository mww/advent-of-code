package day07

import "theelements.org/advent-of-code/common"

func P1() {

}

func P2() {

}

type Node struct {
	Name     string
	Size     int
	Children []*Node
	Parent   *Node
}

func Parse(f string) *Node {
	lines := common.ReadFile(f)
	for _, l := range lines {

	}
}
