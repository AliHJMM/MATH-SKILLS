package functions

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadNumbersFromFile(filename string) ([]int, error) {
	var array []int
	readingFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer readingFile.Close()

	file := bufio.NewScanner(readingFile)

	for file.Scan() {
		line := strings.TrimSpace(file.Text())
		if line == "" {
			continue // Skip empty lines
		}
		if !isValidNumber(line) {
			log.Printf("Invalid number detected: '%s'", line)
			continue // Skip lines with invalid numbers
		}
		parsedValue, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("Error parsing line '%s': %s", line, err)
			continue // Skip lines that can't be parsed
		}
		array = append(array, parsedValue)
	}

	return array, nil
}
