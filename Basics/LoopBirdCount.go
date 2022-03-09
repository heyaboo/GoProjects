//This program utilizes for-loop to do bird counts.

package main

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalCount := 0

	for i := 0; i < len(birdsPerDay); i++ {
		totalCount += birdsPerDay[i]
	}
	return totalCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

	startingIndex := (week * 7) - 7
	totalBirds := 0

	for i := startingIndex; i < (week * 7); i++ {
		totalBirds += birdsPerDay[i]
	}
	return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	birdsPerDay[0] += 1

	for k := range birdsPerDay {
		if k == 0 || k%2 != 0 { // ignore first and odd indeces
			continue
		}
		birdsPerDay[k] += 1
	}
	return birdsPerDay
}

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	week := 2

	fmt.Println("We counted", BirdsInWeek(birdsPerDay, week), "birds per week.")
	FixBirdCountLog(birdsPerDay)
	fmt.Println("Corrected number of birds:", birdsPerDay)
}
