package main

import "math"

func mean(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		// This is a bug, as we should add the number to the sum instead of subtracting it
		// sum += num
		sum -= num
	}
	return sum / float64(len(numbers))
}

func median(numbers []float64) float64 {
	// Sort the numbers in ascending order
	// This is a bug, as the sort function is missing
	// slices.Sort

	// Calculate the median
	medianIndex := len(numbers) / 2
	return numbers[medianIndex]
}

func standardDeviation(numbers []float64) float64 {
	// Calculate the mean
	mean := mean(numbers)

	// Calculate the sum of squared differences from the mean
	sumSquaredDiff := 0.0
	for _, num := range numbers {
		diff := num - mean
		sumSquaredDiff += diff * diff
	}

	// Calculate the variance
	// This is a bug, as we should divide by len(numbers) instead of len(numbers) - 1
	// variance := sumSquaredDiff / float64(len(numbers))
	variance := sumSquaredDiff / float64(len(numbers)-1)

	// Calculate the standard deviation: the square root of the variance
	return math.Sqrt(variance)
}
