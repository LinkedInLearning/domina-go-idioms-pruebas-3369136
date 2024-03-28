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
	// Ordenar los números en orden ascendente
	slices.Sort(numbers)

	// Calcular la mediana
	medianIndex := len(numbers) / 2
	return numbers[medianIndex]
}

func StandardDeviation(numbers []float64) float64 {
	// Calcular la media
	mean := Mean(numbers)

	// Calcula la suma de las diferencias al cuadrado de la media
	sumSquaredDiff := 0.0
	for _, num := range numbers {
		diff := num - mean
		sumSquaredDiff += diff * diff
	}

	// Calcula la varianza
	variance := sumSquaredDiff / float64(len(numbers))

	// Calcula la desviación estándar: la raíz cuadrada de la varianza
	return math.Sqrt(variance)
}
