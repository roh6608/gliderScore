package main

import "testing"

// A test of the minValue function
func TestMinValue(t *testing.T) {

	testData := []struct {
		x []float64
		y float64
	}{
		{[]float64{100.0, 50.0, 25.0}, 25.0},
		{[]float64{50.0, 25.0, 100.0}, 25.0},
		{[]float64{25.0, 100.0, 50.0}, 25.0},
	}

	for _, testTable := range testData {
		testGot := minValue(testTable.x[0], testTable.x[1], testTable.x[2])

		if testGot != 25.0 {
			t.Errorf("minValue was incorrect, expected: %f, got: %f", testTable.y, testGot)
		}
	}
}
