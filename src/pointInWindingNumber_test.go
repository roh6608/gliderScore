package main

import "testing"

// A test of the point in polygon function
func TestPointInWindingNumber(t *testing.T) {

	testData := []struct {
		point    point
		vertices vertices
		output   bool
	}{
		{point{1, 1}, vertices{[]float64{0, 5, 0, 5, 0}, []float64{0, 0, 5, 5, 0}}, true},
		{point{10, 10}, vertices{[]float64{0, 5, 0, 5, 0}, []float64{0, 0, 5, 5, 0}}, false},
	}

	for _, testTable := range testData {
		testGot := pointInWindingNumber(testTable.point, testTable.vertices, 4)

		if testGot != testTable.output {
			t.Errorf("pointInCrossingNumber is incorrect, expected, %t, got %t", testTable.output, testGot)
		}
	}
}

func BenchmarkPointInWindingNumber(b *testing.B) {

	i := 0
	for i < b.N {
		pointInWindingNumber(point{1, 1}, vertices{[]float64{0, 5, 0, 5, 0}, []float64{0, 0, 5, 5, 0}}, 4)
		i += 1

	}
}
