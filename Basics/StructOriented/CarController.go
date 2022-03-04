//This program stimulates a remote car controller. I am learning Golang's struct and this is my first project on it.

package main

import "fmt"

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
	}
	return car
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	track := Track{
		distance: distance,
	}
	return track
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {

	car.distance = car.speed
	car.battery -= car.batteryDrain

	return car
}

//CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	totalBatteryUsage := (track.distance / car.speed) * car.batteryDrain
	var finishing bool

	if totalBatteryUsage <= car.battery {
		finishing = true
	} else if car.battery <= car.batteryDrain || car.battery < 1 {
		finishing = false
	} else {
		finishing = false
	}
	return finishing
}

func main() {
	speed := 5
	batteryDrain := 5
	car := NewCar(speed, batteryDrain)

	distance := 100
	raceTrack := NewTrack(distance)

	CanFinish(car, raceTrack)

	if CanFinish(car, raceTrack) {
		fmt.Println("It will finish this race!")
	} else {
		fmt.Println("Not enough juice in it! Unfortunate~")
	}
}
