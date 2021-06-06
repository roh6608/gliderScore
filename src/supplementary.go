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

// A function to test if a point lies within a polygon using the winding number method, from https://geomalgorithms.com/index.html
// Copyright 2001, 2012, 2021 Dan Sunday
// This code may be freely used and modified for any purpose
// providing that this copyright notice is included with it.
// There is no warranty for this code, and the author of it cannot
// be held liable for any real or imagined damage from its use.
// Users of this code must verify correctness for their application.

type vertices struct {
	x, y []float64
}

type point struct {
	x, y float64
}

func isLeft(P0 point, P1 point, P2 point) float64 {

	return (P1.x-P0.x)*(P2.y-P0.y) - (P2.x-P0.x)*(P1.y-P0.y)
}

func pointInWindingNumber(P point, V vertices, n int) bool {

	wn := 0 // winding number

	i := 0

	for i < n {
		if V.y[i] <= P.y {
			if V.y[i+1] > P.y {
				P0 := point{V.x[i], V.y[i]}
				P1 := point{V.x[i+1], V.y[i+1]}
				if isLeft(P0, P1, P) > 0 {
					wn++
				}
			}
		} else {
			if V.y[i+1] <= P.y {
				P0 := point{V.x[i], V.y[i]}
				P1 := point{V.x[i+1], V.y[i+1]}
				if isLeft(P0, P1, P) < 0 {
					wn--
				}
			}
		}

		i++
	}

	if wn != 0 {
		return true
	} else {
		return false
	}

}

// A function to find if a point lies within a point on the WGS-84 spheroid
func pointInCircle(lat float64, lon float64, centre [][]float64, radius float64) bool {

	distance := vincentyDistance(lat, lon, centre[0][0], centre[0][1])

	if distance <= radius {
		return true
	} else {
		return false
	}
}

// Consider pointers for spped and efficiency
