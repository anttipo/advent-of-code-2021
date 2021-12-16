package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/anttipo/advent-of-code-2021/inputs"
)

func main() {
	filepath := "input"
	part1(filepath)
	part2(filepath)
}

func part1(filepath string) {
	i := inputs.ReadStrings(filepath)

	g := "" // gamma
	e := "" // epsilon

	for k := range i[0] {
		total := 0
		for _, j := range i {
			kek, _ := strconv.Atoi(string(j[k]))
			total = total + kek
		}
		if len(i)/2 < total {
			g += "1"
			e += "0"
		} else {
			g += "0"
			e += "1"
		}
	}
	eresult, _ := strconv.ParseInt(e, 2, 64)
	gresult, _ := strconv.ParseInt(g, 2, 64)
	fmt.Println(eresult * gresult)
}

func part2(filepath string) {

	input, _ := os.Open(filepath)
	defer input.Close()
	sc := bufio.NewScanner(input)

	var oxygenRatingValues []string
	for sc.Scan() {
		oxygenRatingValues = append(oxygenRatingValues, sc.Text())
	}
	co2RatingValues := append([]string{}, oxygenRatingValues...)

	gamma, _ := strconv.ParseInt(filterValues(oxygenRatingValues, 0, true), 2, 64)
	epsilon, _ := strconv.ParseInt(filterValues(co2RatingValues, 0, false), 2, 64)
	fmt.Println(gamma * epsilon)
}

func filterValues(values []string, bitToConsider int, mostCommon bool) string {
	if len(values) == 1 {
		return values[0]
	}

	var valueWithZero, valueWithOne []string

	for _, value := range values {
		if rune(value[bitToConsider]) == '0' {
			valueWithZero = append(valueWithZero, value)
		} else {
			valueWithOne = append(valueWithOne, value)
		}
	}

	if len(valueWithOne) >= len(valueWithZero) == mostCommon {
		return filterValues(valueWithOne, bitToConsider+1, mostCommon)
	}
	return filterValues(valueWithZero, bitToConsider+1, mostCommon)
}
