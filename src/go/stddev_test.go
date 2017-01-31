package main

import (
	"testing"
)

// standard deviations
var stddevTests = []struct {
	points []float64
	answer float64
}{
	{[]float64{5, 10, 10, 10, 15}, 3.2},
	{[]float64{11, 12, 15, 14, 13, 14}, 1.3 },
}

func TestCalcStdDev(t *testing.T) {
	for _, test := range (stddevTests) {
		var res float64 = calcStdDev(test.points)
		if res != test.answer {
			t.Error("Standard Deviation ", res, " != expected ", test.answer);
		}
	}
}

// Means
var meanTests = []struct {
	points []float64
	answer float64
}{
	{[]float64{0.1, 0.2, 0.2, 0.2, 0.3}, 0.2},
	{[]float64{2, 4}, 3},
}

func TestCalcMean(t *testing.T) {
	for _, test := range (meanTests) {
		var res float64 = calcMean(test.points)
		if res != test.answer {
			t.Error("Value ", res, " != expected ", test.answer);
		}
	}
}

var roundTests = []struct {
	input  float64
	result float64
}{
	{0.0003, 0.0},
	{3.4423, 3.4},
	{4.49, 4.5},
	{5.0, 5.0},
	{5.50, 5.5},
	{5.5123, 5.5},
}

func TestRound(t *testing.T) {
	for _, test := range (roundTests) {
		var res float64 = round(test.input, 0.5, 1)
		if res != test.result {
			t.Error("Value ", res, " != expected ", test.result);
		}
	}
}