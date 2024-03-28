package main

import (
	"fmt"
)

func ExampleMean() {
	numbers := []float64{1, 2, 3, 4, 5}
	got := Mean(numbers)
	fmt.Printf("%.2f\n", got)

	// Output:
	// 3.00
}

func ExampleMedian() {
	numbers := []float64{7, 3, 1, 5, 4, 7, 2}
	got := Median(numbers)
	fmt.Printf("%.2f\n", got)

	// Output:
	// 4.00
}

func ExampleStandardDeviation() {
	numbers := []float64{1, 2, 3, 4, 5}
	got := StandardDeviation(numbers)
	fmt.Println(got)

	// Output:
	// 1.4142135623730951
}

func ExampleStandardDeviation_moreNumbers() {
	numbers := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := StandardDeviation(numbers)
	fmt.Println(got)

	// Output:
	// 2.581988897471611
}
