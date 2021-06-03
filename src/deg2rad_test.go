package main

import (
	"math"
	"testing"
)

func TestDeg2Rad(t *testing.T) {
	testData := []struct {
		x float64
		y float64
	}{
		{180, math.Pi},
		{90, math.Pi / 2},
		{60, math.Pi / 3},
		{45, math.Pi / 4},
		{30, math.Pi / 6},
		{15, math.Pi / 12},
		{0, 0},
	}

	for _, testTable := range testData {
		testGot := deg2rad(testTable.x)

		if testGot != testTable.y {
			t.Errorf("deg2rad was incorrect, expected: %f, got: %f", testTable.y, testGot)
		}
	}
}
