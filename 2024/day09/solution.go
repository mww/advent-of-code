package day09

import (
	"log"

	"theelements.org/advent-of-code/common"
)

func P1() {
	input := common.ReadFile("./2024/day09/input.txt")
	data := makeDataSlice(input[0])
	checksum := P1Solution(data)
	log.Printf("The checksum is: %d", checksum)
}

func makeDataSlice(input string) []int {
	data := make([]int, 0, 1024)
	id := 0
	for i, r := range input {
		n := common.Atoi(string(r))
		v := -1 // empty space
		if i%2 == 0 {
			v = id
			id++
		}
		for x := 0; x < n; x++ {
			data = append(data, v)
		}
	}

	return data
}

func calculateChecksum(data []int) int64 {
	cs := int64(0)
	for i := 0; i < len(data); i++ {
		if data[i] == -1 {
			continue
		}
		cs += int64(i * data[i])
	}
	return cs
}

// w and h are the max width and heights so we can determine
// if an antinode is outside of the map
func P1Solution(data []int) int64 {
	head := 0
	tail := len(data) - 1

	for head < tail {
		if data[head] != -1 {
			head++
		} else if data[tail] == -1 {
			tail--
		} else {
			// head is at an empty spot, tail is a point where there is data, swap them
			data[head] = data[tail]
			data[tail] = -1
			//log.Printf("%v", data)
		}
	}

	return calculateChecksum(data)
}

func P2() {
	input := common.ReadFile("./2024/day09/input.txt")
	data := makeDataSlice(input[0])
	checksum := P2Solution(data)
	log.Printf("The checksum is: %d", checksum)
}

func P2Solution(data []int) int64 {
	empty := make([]emptyBlocks, 0, 512)
	for i := 0; i < len(data); i++ {
		if data[i] == -1 {
			block := emptyBlocks{
				start: i,
				size:  1,
			}
			i++

			for i < len(data) && data[i] == -1 {
				block.size++
				i++
			}
			empty = append(empty, block)
		}
	}
	// log.Printf("empty blocks: %v", empty)

	tail := len(data) - 1
	for tail >= 0 {
		if data[tail] != -1 {
			target := data[tail]
			length := 1
			tail--

			for tail >= 0 && data[tail] == target {
				length++
				tail--
			}

			//log.Printf("blocks: target: %d, start index: %d, length: %d", target, tail+1, length)

			// Now see if there are any empty blocks that can fit the data
			for i := 0; i < len(empty); i++ {
				if empty[i].size >= length && empty[i].start < tail+1 {
					for j := 0; j < length; j++ {
						data[empty[i].start+j] = target
						data[tail+1+j] = -1
					}

					empty[i].size -= length
					empty[i].start += length
					//log.Printf("data: %v", data)
					//log.Printf("empty: %v", empty)
					break
				}
			}
		} else {
			tail--
		}
	}
	//log.Printf("end data: %v", data)
	return calculateChecksum(data)
}

type emptyBlocks struct {
	start, size int
}
