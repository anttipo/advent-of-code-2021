package inputs

import (
	"bufio"
	"log"
	"os"
)

func readStrings(filepath string) []string {
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
