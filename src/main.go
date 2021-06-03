package main

import (
	"fmt"
	"math"
)

func maxDaily(d0 float64, dD float64, t0 float64, tT float64) float64 {

	return minValue(1000.0, (1250.0*(d0/dD))-250, (1200.0*(t0/tT))-200)
}

func distancePoints(pointsMaximum float64, nV float64, N float64, rD float64) float64 {

	return pointsMaximum * ((1 - (2*nV)/3*N) * rD)
}

func speedPoints(pointsMaximum float64, rV float64, nV float64, N float64) (speedPoints float64) {

	speedPoints = pointsMaximum * (2 * (rV - 2/3) * nV / N)

	if speedPoints < 0 {
		return 0
	} else {
		return speedPoints
	}
}

func finisherPoints(pointsDistance float64, pointsSpeed float64) float64 {

	return pointsDistance + pointsSpeed
}

func nonFinisherPoints(pointsMaximum float64, rD float64) float64 {

	return pointsMaximum * rD
}

func dayFactor(nD float64, N float64) (dayFactor float64) {

	dayFactor = 1.25 * nD / N

	if dayFactor > 1 {
		return 1
	} else {
		return dayFactor
	}

}

func correctedPoints(dayFactor float64, points float64) float64 {

	return dayFactor * points
}

// A function to find the distance between two points on the WGS-84 ellipsoid, note input is in decimal degrees and output is in metres.
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

// write a test and benchmark for vincenty, also write benchmarks in for the rest of the tests, because it doesnt look to hard

func main() {

	fmt.Println(vincentyDistance(-45, 150, 50, 170))
}

// have a config file where they can input all the necessary information for what they want in the output. or try and use structs that then when they flag it pulls out certain values
