package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	array := []int{}
	total := 0.0

	readingFile, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("Error: Opening File Failure %s", err)
	}
	defer readingFile.Close()

	file := bufio.NewScanner(readingFile)

	for file.Scan() {
		line := file.Text()
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
		total += float64(parsedValue)
	}

	sort.Ints(array)

	fmt.Print("Average: ")
	average := math.Round(total / float64(len(array)))
	fmt.Println(average)

	fmt.Print("Median: ")
	median := 0
	if len(array)%2 == 0 {
		median = int(math.Ceil(float64((array[(len(array)/2)-1])+(array[(len(array)/2)])) / 2))
	} else {
		median = int(math.Ceil(float64(array[len(array)/2])))
	}
	fmt.Println(median)

	fmt.Print("Variance: ")
	standardDeviation := 0.0
	for i := 0; i < len(array); i++ {
		standardDeviation += math.Pow((float64(array[i]) - average), 2)
	}
	variance := math.Round(standardDeviation / float64(len(array)))
	fmt.Println(int(variance))

	fmt.Print("Standard Deviation: ")
	standardDeviation = math.Round(math.Sqrt(standardDeviation / float64(len(array))))
	fmt.Println(standardDeviation)
}

// isValidNumber checks if the input string represents a valid number (no non-numeric characters).
func isValidNumber(s string) bool {
	match, _ := regexp.MatchString(`^\d+$`, s)
	return match
}
