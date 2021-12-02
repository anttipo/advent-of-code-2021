package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anttipo/advent-of-code-2021/inputs"
)

func main() {
	filepath := "input"
	part1(filepath)
	part2(filepath)
}

func part1(filepath string) {
	i := inputs.ReadStrings(filepath)

	horizontal := 0
	depth := 0

	for _, k := range i {
		row := strings.Split(k, " ")
		y, _ := strconv.Atoi(row[1])

		switch {
		case row[0] == "forward":
			horizontal = horizontal + y
		case row[0] == "down":
			depth = depth + y
		case row[0] == "up":
			depth = depth - y
		}
	}
	fmt.Println("Part 1 -- Horizontal: ", horizontal, " and depth: ", depth, " for a total of: ", (horizontal * depth))
}

func part2(filepath string) {
	i := inputs.ReadStrings(filepath)

	horizontal := 0
	depth := 0
	aim := 0

	for _, k := range i {
		row := strings.Split(k, " ")
		y, _ := strconv.Atoi(row[1])

		switch {
		case row[0] == "forward":
			horizontal = horizontal + y
			if aim == 0 {
				// foo
			} else {
				depth = depth + (y * aim)
			}
		case row[0] == "down":
			aim = aim + y
		case row[0] == "up":
			aim = aim - y
		}
	}
	fmt.Println("Part 2 -- Horizontal: ", horizontal, " and depth: ", depth, " for a total of: ", (horizontal * depth))
}
