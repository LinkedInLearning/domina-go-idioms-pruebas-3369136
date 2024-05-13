package main

import (
	"testing"
)

func TestMean(t *testing.T) {
	numbers := []float64{1, 2, 3, 4, 5}
	got := mean(numbers)
	want := 3.0

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestMedian(t *testing.T) {
	numbers := []float64{7, 3, 1, 5, 4, 7, 2}
	got := median(numbers)
	want := 4.0

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestStandardDeviation(t *testing.T) {
	numbers := []float64{1, 2, 3, 4, 5}
	got := standardDeviation(numbers)
	want := 1.4142135623730951

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
