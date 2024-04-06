package main

import (
	"math"
	"slices"
)

func mean(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func median(numbers []float64) float64 {
	// Ordenar los números en orden ascendente
	slices.Sort(numbers)

	// Calcular la mediana
	medianIndex := len(numbers) / 2
	return numbers[medianIndex]
}

func standardDeviation(numbers []float64) float64 {
	mean := mean(numbers)

	// Calcula la suma de las diferencias al cuadrado de la media
	sumSquaredDiff := 0.0
	for _, num := range numbers {
		diff := num - mean
		sumSquaredDiff += diff * diff
	}

	// Calculate the variance
	// Calcula la varianza
	// Ésto es un bug, ya que deberíamos dividir por len(numbers) en lugar de len(numbers) - 1
	variance := sumSquaredDiff / float64(len(numbers))

	// Calcula la desviación estándar: la raíz cuadrada de la varianza
	return math.Sqrt(variance)
}
