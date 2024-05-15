package main

import (
	"fmt"
	"math"
)
func Sphere() func() float64 {
	var radius = 0.0
	result := func() float64 {
		radius += 1.0
		volume := 4 / 3 * math.Pi * radius * radius * radius
		return volume
	}

	return result

}

// Main Function
func main() {

	sVol := Sphere() // calling outer func

	// calling inner function ehich holds the outer func's variable
	fmt.Println("Volume of Sphere is:", sVol())
	fmt.Println("Volume of Sphere is:", sVol())
}
