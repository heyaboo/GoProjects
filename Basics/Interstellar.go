package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Println("Fuel level:", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Println("Current location:", planet)
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("Insufficient fuel to fly to destination.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {

	fuelRemaining := fuel
	fuelCost := calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {

	defer fmt.Println("It was a fun journey!") //defer prints this line at the end of program

	//Create and assign required variables
	fuel := 1000000
	var planetChoice string
	fmt.Print("Where would you like to fly to? ")
	fmt.Scan(&planetChoice)

	//Testing the mini program
	fuelGauge(fuel)
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)

}

// Challenge portion to be added:
// Add a variable that keeps track of which planet your spaceship is on.
// Create a function that returns your spaceship back to your home planet.
// Add more destinations and allow for traveling between different planets.
