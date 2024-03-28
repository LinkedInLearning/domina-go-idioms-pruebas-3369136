package main

import (
	"log"
	"math/rand"
)

func main() {
	var numbers []float64
	for i := 0; i < 100; i++ {
		ran := rand.Float64() * 100
		numbers = append(numbers, float64(i)*ran)
	}

	// Calculate the mean of the numbers
	log.Println("Mean:", Mean(numbers))

	// Calculate the median of the numbers
	log.Println("Median:", Median(numbers))

	// Calculate the standard deviation of the numbers
	log.Println("Standard Deviation:", StandardDeviation(numbers))
}
