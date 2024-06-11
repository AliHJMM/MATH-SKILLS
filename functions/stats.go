package functions

import (
	"math"
	"sort"
)

func CalculateStatistics(array []int) (average, median, variance, stddev float64) {
	sort.Ints(array)
	total := 0.0
	for _, num := range array {
		total += float64(num)
	}

	average = math.Round(total / float64(len(array)))

	if len(array)%2 == 0 {
		median = math.Ceil(float64((array[(len(array)/2)-1])+(array[(len(array)/2)])) / 2)
	} else {
		median = math.Ceil(float64(array[len(array)/2]))
	}

	var varianceSum float64
	for _, num := range array {
		varianceSum += math.Pow((float64(num) - average), 2)
	}
	variance = math.Round(varianceSum / float64(len(array)))
	stddev = math.Round(math.Sqrt(variance))

	return average, median, variance, stddev
}
