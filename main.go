package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	array, total, median, standardDeviation, average := []int{}, 0.0, 0, 0.0, 0.0
	readingFile, err := os.Open("data.txt")
	defer readingFile.Close()

	if err != nil {
		log.Fatalf("Error: Opening File Failure %s", err)
	}

	file := bufio.NewScanner(readingFile)
	file.Split(bufio.ScanLines)

	for file.Scan() {
		parsedValue, err := strconv.Atoi(file.Text())
		if err != nil {
			fmt.Printf("%T \n %v", parsedValue, parsedValue)
		}
		array = append(array, parsedValue)
		total += float64(parsedValue)
	}
	sort.Ints(array)

	fmt.Print("Average: ")
	average = math.Round(total / float64(len(array)))
	fmt.Println(average)

	fmt.Print("Median: ")
	if len(array)%2 == 0 {
		median = int(math.Ceil(float64((array[(len(array)/2)-1])+(array[(len(array)/2)])) / 2))
	} else {
		median = int(math.Ceil(float64(array[(len(array) / 2)])))
	}

	fmt.Println(median)
	fmt.Print("Variance: ")
	for i := 0; i < len(array); i++{
		standardDeviation += math.Pow((float64(array[i])- average), 2)
	}
	variance := math.Round(standardDeviation / float64(len(array)))
	fmt.Println(int(variance))

}
