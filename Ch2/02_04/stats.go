package main

import (
	"math"
	"slices"
)

func Mean(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func Median(numbers []float64) float64 {
	// Sort the numbers in ascending order
	slices.Sort(numbers)

	// Calculate the median
	medianIndex := len(numbers) / 2
	return numbers[medianIndex]
}

func StandardDeviation(numbers []float64) float64 {
	// Calculate the mean
	mean := Mean(numbers)

	// Calculate the sum of squared differences from the mean
	sumSquaredDiff := 0.0
	for _, num := range numbers {
		diff := num - mean
		sumSquaredDiff += diff * diff
	}

	// Calculate the variance
	variance := sumSquaredDiff / float64(len(numbers))

	// Calculate the standard deviation: the square root of the variance
	return math.Sqrt(variance)
}
