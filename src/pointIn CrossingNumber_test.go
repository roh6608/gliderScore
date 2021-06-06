package main

import "testing"

func TestPointInCrossingNumber(t *testing.T) {

	testData := []struct {
		x        float64
		y        float64
		vertices [][]float64
		output   int
	}{
		{2.5, 2.5, [][]float64{{0, 0}, {5, 0}, {5, 5}, {0, 5}, {0, 0}}, 1},
		{10, 10, [][]float64{{0, 0}, {5, 0}, {5, 5}, {0, 5}, {0, 0}}, 0},
	}

	for _, testTable := range testData {
		testGot := pointInCrossingNumber(testTable.y, testTable.x, testTable.vertices)

		if testGot != testTable.output {
			t.Errorf("pointInCrossingNumber is incorrect, expected, %d, got %d", testTable.output, testGot)
		}
	}
}
