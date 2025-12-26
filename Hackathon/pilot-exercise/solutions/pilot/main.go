package main

import "fmt"

// Pilot struct definition
// Fields match the assignments in main function
type Pilot struct {
	Name     string  // Pilot's name
	Life     float64 // Life points
	Age      int     // Pilot's age
	Aircraft uint    // Aircraft ID (matches const type)
}

func main() {
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1
	fmt.Println(donnie)
}

const AIRCRAFT1 = 1
