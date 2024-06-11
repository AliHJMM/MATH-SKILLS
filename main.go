package main

import (
	"MATH-SKILLS/functions"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}

	filename := os.Args[1]

	array, err := functions.ReadNumbersFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading from file: %s", err)
	}

	if len(array) == 0 {
		log.Println("No valid numbers to process.")
		return
	}

	average, median, variance, stddev := functions.CalculateStatistics(array)

	fmt.Printf("Average: %.0f\n", average)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard Deviation: %.0f\n", stddev)
}
