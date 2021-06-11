package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

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

// consider pointers for efficiency

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

// A function to find if a point lies within a circle on the WGS-84 ellipsoid
func pointInCircle(lat float64, lon float64, centre [][]float64, radius float64) bool {

	distance := vincentyDistance(lat, lon, centre[0][0], centre[0][1])

	if distance <= radius {
		return true
	} else {
		return false
	}
}

/* A function to find the distance between two points on the WGS-84 ellipsoid using the Vincenty algorithm, note input is in decimal degrees and output is in metres.

   Equations from https://en.wikipedia.org/wiki/Vincenty%27s_formulae, implementation a derivative of implementation found at
   https://www.johndcook.com/blog/2018/11/24/spheroid-distance/ as per https://www.fai.org/page/world-distance-calculator.
*/
func vincentyDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) (distance float64) {

	lat1 = deg2rad(lat1)
	lat2 = deg2rad(lat2)
	lon1 = deg2rad(lon1)
	lon2 = deg2rad(lon2)

	a := 6378137.0
	f := 1 / 298.257223563
	b := (1 - f) * a

	tolerance := 1e-12

	phi1, phi2 := lat1, lat2
	U1 := math.Atan((1 - f) * math.Tan(phi1))
	U2 := math.Atan((1 - f) * math.Tan(phi2))
	L1, L2 := lat1, lat2
	L := L2 - L1

	lambdaOld := L + 0

	for {
		t := math.Pow(math.Cos(U2)*math.Sin(lambdaOld), 2)
		t += math.Pow((math.Cos(U1)*math.Sin(U2) - math.Sin(U1)*math.Cos(U2)*math.Cos(lambdaOld)), 2)
		sinSigma := math.Pow(t, 0.5)
		cosSigma := math.Sin(U1)*math.Sin(U2) + math.Cos(U1)*math.Cos(U2)*math.Cos(lambdaOld)
		sigma := math.Atan2(sinSigma, cosSigma)

		sinAlpha := math.Cos(U1) * math.Cos(U2) * math.Sin(lambdaOld) / sinSigma
		cosSqAlpha := 1 - math.Pow(sinAlpha, 2)
		cos2SigmaM := cosSigma - 2*math.Sin(U1)*math.Sin(U2)/cosSqAlpha
		c := f * cosSqAlpha * (4 + f*(4-3*cosSqAlpha)) / 16

		t = sigma + c*sinSigma*(cos2SigmaM+c*cosSigma*(-1+2*math.Pow(cos2SigmaM, 2)))
		lambdaNew := L + (1-c)*f*sinAlpha*t

		if math.Abs(lambdaNew-lambdaOld) <= tolerance {
			break
		} else {
			lambdaOld = lambdaNew
		}

		u2 := cosSqAlpha * ((math.Pow(a, 2) - math.Pow(b, 2)) / math.Pow(b, 2))
		A := 1 + (u2/16384)*(4096+u2*(-768+u2*(320-175*u2)))
		B := (u2 / 1024) * (256 + u2*(-128+u2*(74-47*u2)))
		t = cos2SigmaM + 0.25*B*(cosSigma*(-1+2*math.Pow(cos2SigmaM, 2)))
		t -= (B / 6) * cos2SigmaM * (-3 + 4*math.Pow(sinSigma, 2)) * (-3 + 4*math.Pow(cos2SigmaM, 2))
		deltaSigma := B * sinSigma * t
		distance = b * A * (sigma - deltaSigma)
	}

	return distance
}

func readFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

type bRecord struct {
	time             []time.Time
	latitude         []float64
	longitude        []float64
	fixValidity      []string
	pressureAltitude []float64
	gnssAltitude     []float64
}

// make the b recod parser only use the one struct, should be possible
func parseB(file []string) bRecord {
	// implement a check on the I records to see if any additionals have to be parsed also add these to the structs
	// can do this implementation without the two different structs, implement it this way
	// write an I record parser, that can parse the I record and then with that then know how to parse the B record
	// so Should write the I record parser first
	// write out an I record struct that has all the possibilities that it can have and then it records the bytes
	// in which it is located, this can then be used in the B record parser, then use if statements to check what is in the
	// I record parser

	var record bRecord
	for i := 0; i < len(file); i++ {
		if string(file[i][0]) == "B" {
			time, _ := time.Parse("150405", file[i][1:7])
			record.time = append(record.time, time)

			// checking if latitude is north or south, and parsing as needed
			// need to check charcter index for the latitude, think they are fine but need to check
			if string(file[i][14]) == "N" {
				degrees1, _ := strconv.ParseFloat(file[i][7:9], 64)
				degrees2, _ := strconv.ParseFloat(file[i][9:11]+"."+file[i][11:14], 64)
				record.latitude = append(record.latitude, (degrees1 + degrees2/60))
			} else {
				degrees1, _ := strconv.ParseFloat(file[i][7:9], 64)
				degrees2, _ := strconv.ParseFloat(file[i][9:11]+"."+file[i][11:14], 64)
				record.latitude = append(record.latitude, -(degrees1 + degrees2/60))
			}

			// checking if longitude is east or west, and parsing as needed
			// indexing string is incorrect here, not quite sure where yet.
			if string(file[i][23]) == "E" {
				degrees1, _ := strconv.ParseFloat(file[i][15:18], 64)
				degrees2, _ := strconv.ParseFloat(file[i][18:20]+"."+file[i][20:23], 64)
				record.longitude = append(record.longitude, (degrees1 + degrees2/60))
			} else {
				degrees1, _ := strconv.ParseFloat(file[i][15:18], 64)
				degrees2, _ := strconv.ParseFloat(file[i][18:20]+"."+file[i][20:23], 64)
				record.longitude = append(record.longitude, -(degrees1 + degrees2/60))
			}

			record.fixValidity = append(record.fixValidity, file[i][24:25])
			pressureAltitude, _ := strconv.ParseFloat(file[i][25:30], 64)
			record.pressureAltitude = append(record.pressureAltitude, pressureAltitude)
			gnnsAltitude, _ := strconv.ParseFloat(file[i][30:35], 64)
			record.gnssAltitude = append(record.gnssAltitude, gnnsAltitude)
		}
	}

	return record

}
