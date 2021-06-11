package main

// Function to return the maximum daily points available, as per rule 44.2
func maxDaily(d0 float64, dD float64, t0 float64, tT float64) float64 {

	return minValue(1000.0, (1250.0*(d0/dD))-250, (1200.0*(t0/tT))-200)
}

// Function to return the value of distance points, as per rule 44.3
func distancePoints(pointsMaximum float64, nV float64, N float64, rD float64) float64 {

	return pointsMaximum * ((1 - (2*nV)/3*N) * rD)
}

// Function to return the value of the speed points, as per rule 44.3
func speedPoints(pointsMaximum float64, rV float64, nV float64, N float64) (speedPoints float64) {

	speedPoints = pointsMaximum * (2 * (rV - 2/3) * nV / N)

	if speedPoints < 0 {
		return 0
	} else {
		return speedPoints
	}
}

// Function to return the value of the points for a task with finishers, as per rule 44.3
func finisherPoints(pointsDistance float64, pointsSpeed float64) float64 {

	return pointsDistance + pointsSpeed
}

// Function to return the value of points for a task with no finishers, as per rule 44.4
func nonFinisherPoints(pointsMaximum float64, rD float64) float64 {

	return pointsMaximum * rD
}

// Function to return the value of the Day Factor f, as per rule 44.6
func dayFactor(nD float64, N float64) (dayFactor float64) {

	dayFactor = 1.25 * nD / N

	if dayFactor > 1 {
		return 1
	} else {
		return dayFactor
	}

}

// Function to return the value of the Corrected Points Pc, as per rule 44.6
func correctedPoints(dayFactor float64, points float64) float64 {

	return dayFactor * points
}

func main() {

}
