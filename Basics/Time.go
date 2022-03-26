//This program helps to understand how the time package works. Source of exercise: Exercism

package main

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	timeParsed, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil { //error handling
		panic(err)
	}
	return timeParsed
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	targetDate := Schedule(date)

  return targetDate.Before(time.Now())  //Others like targetTime.After() and it will show if the time compared happens after a given time.
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	timeParsed, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	if timeParsed.Hour() >= 12 && timeParsed.Hour() < 18 {  //We can extract hours, minutes, seconds, months and years by using .Year() etc.
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	timeParsed, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("You have an appointment on %s", timeParsed.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	t, _ := time.Parse("2006-01-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return t
}

func main() {
	appointmentTime := "7/25/2022 13:45:00"

	fmt.Println(Schedule(appointmentTime))
	fmt.Println(HasPassed(appointmentTime))
	fmt.Println(Description(appointmentTime))
	fmt.Println(AnniversaryDate())
}
