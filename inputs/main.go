package inputs

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadStrings(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := make([]string, 0)

	for scanner.Scan() {
		i := scanner.Text()
		if err != nil {
			continue
		}
		data = append(data, i)
	}

	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ReadInts(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := make([]int, 0)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		data = append(data, i)
	}

	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ReadBytes(filepath string) []byte {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := make([]byte, 0)

	for scanner.Scan() {
		i := scanner.Text()
		if err != nil {
			continue
		}
		data = append(data, i...)
	}

	if err != nil {
		log.Fatal(err)
	}
	return data
}
