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

	// Calcula la media de los números
	log.Println("Mean:", mean(numbers))

	// Calcula la mediana de los números
	log.Println("Median:", median(numbers))

	// Calcula la desviación estándar de los números
	log.Println("Standard Deviation:", standardDeviation(numbers))
}
