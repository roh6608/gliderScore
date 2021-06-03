package main

import "math"

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

// A function to find the distance between two points on the WGS-84 ellipsoid, note that the input is in radians.
func vincentyDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) (distance float64) {

	a := 6378137.0
	f := 1 / 298.257223563
	b := (1 - f) * a

	tolerance := 1e-11

	phi1, phi2 := lat1, lat2
	U1 := math.Atan((1 - f) * math.Tan(phi1))
	U2 := math.Atan((1 - f) * math.Tan(phi2))
	L1, L2 := lat1, lat2
	L := L2 - L1

	lambdaOld := L + 0

	for {
		t := math.Pow(math.Cos(U2)*math.Sin(lambdaOld), 2)
		t += math.Pow((math.Cos(U1)*math.Sin(U2) - math.Sin(U1)), 2)
	}

	return s
}

func main() {

}

// have a config file where they can input all the necessary information
