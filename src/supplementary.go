package main

import "math"

// A function to return the minimum value of 3 numbers
func minValue(a float64, b float64, c float64) float64 {

	if (a < b) && (a < c) {
		return a
	}

	if (b < a) && (b < c) {
		return b
	} else {
		return c
	}

}

// A function to convert decimal degrees into radians
func deg2rad(deg float64) float64 {

	return deg * math.Pi / 180
}

// A function to test if a point lies within a polygon using the crossing number method, as used from https://geomalgorithms.com/index.html
// Copyright 2001, 2012, 2021 Dan Sunday
// This code may be freely used and modified for any purpose
// providing that this copyright notice is included with it.
// There is no warranty for this code, and the author of it cannot
// be held liable for any real or imagined damage from its use.
// Users of this code must verify correctness for their application.
func pointInCrossingNumber(lat float64, lon float64, vertices [][]float64) int {

	cn := 0

	i := 0

	// array indexing may be backwards here
	for i < len(vertices)-1 {
		if ((vertices[i][1] <= lat) && (vertices[i+1][1] > lat)) || ((vertices[i][1] > lat) && vertices[i+1][1] <= lat) {

			vt := (lat - vertices[i][1]) / (vertices[i+1][1] - vertices[i][1])

			if vertices[i][0] < vertices[i][0]+vt*(vertices[i+1][0]-vertices[i][0]) {
				cn += 1
			}
		}

		i += 1
	}

	return cn & 1

}

// re-write the above function using pointers and structs as in their example

// write function for if it is inside a circle using the math stack exchange example, also check if there is more efficient
// method, area testing algorithms will be used for penalty points and other polygon related things that need to be tested
// write ut both types and then can use tests and benchmarks to see which one operates quicker
// likewise for finding if the point lies within a circle
// when writing it also only compare d^2 with r^2 because it is more efficient
// turn point will use vincenty method and just test that a point falls within the distance, also return a message in a log
// file if points are very close to the boundary
// for turn point radius filter out the points that are not relevant to the turn point so it only has to test a small amount of points
