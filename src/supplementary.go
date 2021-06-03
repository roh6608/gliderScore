package main

import "math"

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

func deg2rad(deg float64) float64 {

	return deg * math.Pi / 180
}
