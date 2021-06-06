package main

import "testing"

func TestPointInCicrle(t *testing.T) {
	testData := []struct {
		lat    float64
		lon    float64
		centre [][]float64
		radius float64
		output bool
	}{
		{-29.5, 150.5, [][]float64{{-30, 150}}, 100000, true},
		{-29.5, 150.5, [][]float64{{-30, 150}}, 50000, false},
	}

	for _, testTable := range testData {
		testGot := pointInCircle(testTable.lat, testTable.lon, testTable.centre, testTable.radius)

		if testGot != testTable.output {
			t.Errorf("pointInCircle is incorrect, expected : %t, got: %t", testTable.output, testGot)
		}
	}
}
