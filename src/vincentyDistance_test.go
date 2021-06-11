package main

import "testing"

func TestVincentyMethod(t *testing.T) {
	// testData y data is from the geoscience australia geodetic calculator
	testData := []struct {
		x []float64
		y float64
	}{
		{[]float64{-30, 150, -29, 151}, 147269.629},
		{[]float64{-30, 150, -29.5, 150.5}, 73558.4},
	}

	for _, testTable := range testData {
		testGot := vincentyDistance(testTable.x[0], testTable.x[1], testTable.x[2], testTable.x[3])

		tolerance := 0.1

		if testGot < testTable.y-tolerance || testGot > testTable.y+tolerance {
			t.Errorf("vincentyDistance was incorrect, expected: %f, got: %f", testTable.y, testGot)
		}
	}
}

func BenchmarkVincenyMethod(b *testing.B) {

	i := 0
	for i < b.N {
		vincentyDistance(-30, 150, -30.01, 150.01)
		i += 1

	}
}

// To run benchmark put go test -bench=.
