package main

import (
	"fmt"

	"github.com/anttipo/advent-of-code-2021/inputs"
)

func main() {
	filepath := "input"

	part1(filepath)
	part2(filepath)
}

func part1(filepath string) {
	i := inputs.ReadInts(filepath)
	count := 0
	previous := i[0]
	for j, k := range i {
		if k > previous {
			count++
		}
		previous = i[j]
	}
	fmt.Println("Day 1 part 1: ", count)
}

func part2(filepath string) {
	i := inputs.ReadInts(filepath)

	count := 0

	for j := 1; j+2 < len(i); j++ {
		prev := i[j-1] + i[j] + i[j+1]
		curr := i[j] + i[j+1] + i[j+2]
		if curr > prev {
			count++
		}
	}
	fmt.Println("Day 1 part 2: ", count)
}
