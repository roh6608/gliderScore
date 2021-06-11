package main

import (
	"fmt"
)

func main() {

	file := readFile("../example/files/crash.igc")
	brec := parseB(file)

	fmt.Println(totalFlightDistance(brec.latitude, brec.longitude))

}
